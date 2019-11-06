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

