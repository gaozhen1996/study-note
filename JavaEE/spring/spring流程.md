# 前言

spring ioc最核心的步骤，是注册配置bean，获取扫描路径来扫描bean。

发现一篇写得非常好的博客 https://www.ibm.com/developerworks/cn/java/j-lo-spring-principle/ 

# 1.概念

- BeanDefinition

  这个类在spring中非常的重要，spring通过BeanDefinition来描述一个类，比如是否是单例的，是否是懒加载的等等

-  DefaultListableBeanFactory 

  存放BeanDefinition的工厂

# 一、注册

## 1.核心类介绍

- AnnotationConfigApplicationContext 

  主要作用：整个spring的环境，是使用spring的入口。

- AnnotatedBeanDefinitionReader

  主要作用：由类名就可知，是被注解的BeanDefinition的读取器，将类转为BeanDefinition

## 2.流程

AnnotationConfigApplicationContext.java

```java
	public AnnotationConfigApplicationContext(Class<?>... annotatedClasses) {
		this();
        //注册加了注解的类
		register(annotatedClasses);
		refresh();
	}

	public void register(Class<?>... annotatedClasses) {
        //通过reader来读取被注解类，这个reader对象就是AnnotatedBeanDefinitionReader的实例
		this.reader.register(annotatedClasses);
	}
```

AnnotatedBeanDefinitionReader.java

```java
	public void register(Class<?>... annotatedClasses) {
        //遍历参数，来注册被注解类
		for (Class<?> annotatedClass : annotatedClasses) {
			registerBean(annotatedClass);
		}
	}

	public void registerBean(Class<?> annotatedClass) {
		doRegisterBean(annotatedClass, null, null, null);
	}

	<T> void doRegisterBean(Class<T> annotatedClass, @Nullable Supplier<T> 											instanceSupplier, @Nullable String name,
							@Nullable Class<? extends Annotation>[] qualifiers, 									BeanDefinitionCustomizer... definitionCustomizers) {
		//new一个BeanDefinition对象
		AnnotatedGenericBeanDefinition abd = new 							                                                 AnnotatedGenericBeanDefinition(annotatedClass);
        //判断是否需要跳过，主要依据是判断这个类是否加了注解
		if (this.conditionEvaluator.shouldSkip(abd.getMetadata())) {
			return;
		}
        //设置实例提供者，这个set的是null
		abd.setInstanceSupplier(instanceSupplier);
		//省略其他设置BeanDefinition属性代码
        //封装BeanDefinition和beanname为BeanDefinitionHolder
		BeanDefinitionHolder definitionHolder = new BeanDefinitionHolder(abd, beanName);
		definitionHolder = AnnotationConfigUtils
            .applyScopedProxyMode(scopeMetadata, definitionHolder, this.registry);
        /**注册到beanFactory中，这里的this.registry的实例就是	                                         AnnotationConfigApplicationContext的实例，
        	而完成注册动作的还是DefaultListableBeanFactory对象
        **/
		BeanDefinitionReaderUtils.registerBeanDefinition(definitionHolder, 		                                                  this.registry);
	}
```



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
        /*委托PostProcessorRegistrationDelegate,来对beanFactory执行BeanFactoryPostProcessor
        值得一提的是，在之前注册的步骤，实例化AnnotatedBeanDefinitionReader对象的时候，
        在构造方法中，调用了
        AnnotationConfigUtils.registerAnnotationConfigProcessors(this.registry)
        这个方法注册了一些默认的后置处理器对象，具体可以看源码
        重要，重要，重要！！！这个时候添加了非常重要的类ConfigurationClassPostProcessor，之后的解析		 全是在这个类中完成的。
        */
		PostProcessorRegistrationDelegate
            .invokeBeanFactoryPostProcessors(beanFactory,getBeanFactoryPostProcessors());
	}
```

PostProcessorRegistrationDelegate.java 

```java
	public static void invokeBeanFactoryPostProcessors(
			ConfigurableListableBeanFactory beanFactory, 
            List<BeanFactoryPostProcessor> beanFactoryPostProcessors) {
        /**
         *值得注意的是，在这个方法中，定义了两个List<BeanDefinitionRegistryPostProcessor>
         *一个是registryProcessors，这个数据来源是beanFactoryPostProcessors，它的作用是获取
         *用户自定义的BeanDefinitionRegistryPostProcessor，来执行
         *而
         */
        List<BeanDefinitionRegistryPostProcessor> registryProcessors = new ArrayList<>();
        for (BeanFactoryPostProcessor postProcessor : beanFactoryPostProcessors) {
            if (postProcessor instanceof BeanDefinitionRegistryPostProcessor) {
                BeanDefinitionRegistryPostProcessor registryProcessor =
                    (BeanDefinitionRegistryPostProcessor) postProcessor;
                /**
                 *执行用户自定义的BeanDefinitionRegistryPostProcessor
                 *测试代码https://github.com/gaozhen1996/study-java/blob/master
                        /src/main/java/com/gz/javastudy/springapp/TestMain.java的
                   testExecuteCustomBeanFactoryPostProcessor();方法
                 */
                registryProcessor.postProcessBeanDefinitionRegistry(registry);
                registryProcessors.add(registryProcessor);
            }
            else {
                regularPostProcessors.add(postProcessor);
            }
		}
        //重要，重要，重要！！！执行BeanDefinitionRegistryPostProcessor
        List<BeanDefinitionRegistryPostProcessor> currentRegistryProcessors 
                                   = new ArrayList<>();
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
        /**第一步：获取容器中所有的配置bean，获取的代码省略
        ，需要注意的是，加了Component，ComponentScan
         Import,ImportResource注解的也会被作为配置类处理
        */
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

ClassPathBeanDefinitionScanner.java

```java
    protected Set<BeanDefinitionHolder> doScan(String... basePackages) {
        for (String basePackage : basePackages) {
            //扫描返回BeanDefinition
            Set<BeanDefinition> candidates = findCandidateComponents(basePackage);
                for (BeanDefinition candidate : candidates) {
        }
    }
    
    public Set<BeanDefinition> findCandidateComponents(String basePackage) {
        if (this.componentsIndex != null && indexSupportsIncludeFilters()) {
            return addCandidateComponentsFromIndex(this.componentsIndex, basePackage);
        }
        else {
             //真正完成扫描的方法
            return scanCandidateComponents(basePackage);
        }
    }
```

