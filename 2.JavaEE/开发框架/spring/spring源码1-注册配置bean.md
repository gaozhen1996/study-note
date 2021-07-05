

# 前言

spring ioc最核心的步骤，是注册配置bean，获取扫描路径来扫描bean。

发现一篇写得非常好的博客 https://www.ibm.com/developerworks/cn/java/j-lo-spring-principle/ 

## 概念

- BeanDefinition

  这个类在spring中非常的重要，spring通过BeanDefinition来描述一个类，比如是否是单例的，是否是懒加载的等等

-  DefaultListableBeanFactory 

  存放BeanDefinition的工厂

# 1.核心类介绍

- AnnotationConfigApplicationContext 

  主要作用：整个spring的环境，是使用spring的入口。

- AnnotatedBeanDefinitionReader

  主要作用：由类名就可知，是被注解的BeanDefinition的读取器，将类转为BeanDefinition

# 2.流程

## AnnotationConfigApplicationContext.java

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

## AnnotatedBeanDefinitionReader.java

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



