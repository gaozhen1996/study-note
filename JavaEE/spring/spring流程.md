# 前言

spring ioc最核心的步骤，是注册配置bean，扫描bean。

发现一篇写得非常好的博客 https://www.ibm.com/developerworks/cn/java/j-lo-spring-principle/ 



# 一、注册



# 二、扫描

## 1.核心类介绍

- AnnotationConfigApplicationContext 

  主要作用：整个spring的环境，是使用spring的入口。

- PostProcessorRegistrationDelegate

  主要作用 : AbstractApplicationContext的后处理器处理的委托，作用是执行beanFactory的BeanFactoryPostProcessors。

- ConfigurationClassPostProcessor 

  主要作用：解析配置类，比如我们在使用spring的使用会在先注册配置类到spring容器中，spring会解析这个这个配置类，这个类的作用就是解析配置类的。

- ClassPathBeanDefinitionScanner

  主要作用：扫描包。将配置类上的@ComponentScan包扫描，返回Set<BeanDefinitionHolder> 。真正扫描的方法是org.springframework.context.annotation.ClassPathScanningCandidateComponentProvider.scanCandidateComponents(String basePackage)，这个类是ClassPathBeanDefinitionScanner的父类

## 2.流程

**说明：这个过程只是关注spring的扫描包的流程，因此简化了展示的代码**

AbstractApplicationContext.java

```java
	@Override
	public void refresh() throws BeansException, IllegalStateException {
		//源码对这个类的注释是：调用上下文中注册为bean的工厂处理器。
		invokeBeanFactoryPostProcessors(beanFactory);
	}
	
	protected void invokeBeanFactoryPostProcessors(
        				ConfigurableListableBeanFactory beanFactory) {
        //委托PostProcessorRegistrationDelegate,来对beanFactory执行BeanFactoryPostProcessor
        //在这个类中，有一个属性List<BeanFactoryPostProcessor>，存放时是需要执行的后置处理器
		PostProcessorRegistrationDelegate
            .invokeBeanFactoryPostProcessors(beanFactory,getBeanFactoryPostProcessors());
	}
```

PostProcessorRegistrationDelegate.java 

```java
	public static void invokeBeanFactoryPostProcessors(
			ConfigurableListableBeanFactory beanFactory, 
            List<BeanFactoryPostProcessor> beanFactoryPostProcessors) {
        //执行BeanDefinitionRegistryPostProcessor
        invokeBeanDefinitionRegistryPostProcessors(currentRegistryProcessors, registry);
    } 


	private static void invokeBeanDefinitionRegistryPostProcessors(
        Collection<? extends BeanDefinitionRegistryPostProcessor> postProcessors, 													BeanDefinitionRegistry registry) {
        //干扰BeanDefinitionRegistry的处理
        for (BeanDefinitionRegistryPostProcessor postProcessor : postProcessors) {
			postProcessor.postProcessBeanDefinitionRegistry(registry);
		}
        
    }
```

ConfigurationClassPostProcessor.java

这个类实现了BeanDefinitionRegistryPostProcessor接口

```java
	public void postProcessBeanDefinitionRegistry(BeanDefinitionRegistry registry) {
		processConfigBeanDefinitions(registry);
	}
	
	public void processConfigBeanDefinitions(BeanDefinitionRegistry registry) {
        //第一步：获取容器中所有的配置bean，获取的代码省略，需要注意的是，加了Component，ComponentScan
        //Import,ImportResource注解的也会被作为配置类处理
        List<BeanDefinitionHolder> configCandidates = new ArrayList<>();
        Set<BeanDefinitionHolder> candidates = new LinkedHashSet<>(configCandidates);
        do {
            //第二步：解析配置bean
            parser.parse(candidates);
        }(!candidates.isEmpty());
    }

	public void parse(Set<BeanDefinitionHolder> configCandidates) {
        for (BeanDefinitionHolder holder : configCandidates) {
			BeanDefinition bd = holder.getBeanDefinition();
            if (bd instanceof AnnotatedBeanDefinition) {
                	//判断如果是AnnotatedBeanDefinition才解析
					parse(((AnnotatedBeanDefinition) bd).getMetadata(), 														holder.getBeanName());
				}
        }
    }

	protected final void parse(AnnotationMetadata metadata, String beanName) throws 																		IOException {
        //封装为ConfigurationClass对象解析
		processConfigurationClass(new ConfigurationClass(metadata, beanName));
	}

    /**
    *解析ConfigurationClass对象
    */
	protected void processConfigurationClass(ConfigurationClass configClass) throws 																		IOException {
        do {
			sourceClass = doProcessConfigurationClass(configClass, sourceClass);
		}
		while (sourceClass != null);
    }

	protected final SourceClass doProcessConfigurationClass(ConfigurationClass 			   					configClass, SourceClass sourceClass)throws IOException {
        // Process any @ComponentScan annotations (源码的注释)
 		Set<AnnotationAttributes> componentScans = AnnotationConfigUtils
     				.attributesForRepeatable(sourceClass.getMetadata(), 			                                                  ComponentScans.class, 
                                             ComponentScan.class);
        
        for (AnnotationAttributes componentScan : componentScans) {
			// config类被@ComponentScan注释->立即执行扫描
			Set<BeanDefinitionHolder> scannedBeanDefinitions =
				this.componentScanParser.parse(componentScan, 				      										sourceClass.getMetadata().getClassName());
        }
        
    }
```

ComponentScanAnnotationParser.java

```java
	public Set<BeanDefinitionHolder> parse(AnnotationAttributes componentScan, final 															String declaringClass) {
        //获取需要扫描的包
        String[] basePackagesArray = componentScan.getStringArray("basePackages");
        return scanner.doScan(StringUtils.toStringArray(basePackages));
    }
```

