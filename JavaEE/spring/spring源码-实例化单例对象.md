# 一、调试堆栈

```java
com.gz.javastudy.springapp.TestSpringApp at localhost:61545	
	Thread [main] (Suspended)	
		owns: ConcurrentHashMap<K,V>  (id=30)	
		owns: Object  (id=31)	
		BeanUtils.instantiateClass(Constructor<T>, Object...) line: 171	
		CglibSubclassingInstantiationStrategy(SimpleInstantiationStrategy).instantiate(RootBeanDefinition, String, BeanFactory) line: 87	
		DefaultListableBeanFactory(AbstractAutowireCapableBeanFactory).instantiateBean(String, RootBeanDefinition) line: 1279	
		DefaultListableBeanFactory(AbstractAutowireCapableBeanFactory).createBeanInstance(String, RootBeanDefinition, Object[]) line: 1181	
		DefaultListableBeanFactory(AbstractAutowireCapableBeanFactory).doCreateBean(String, RootBeanDefinition, Object[]) line: 555	
		DefaultListableBeanFactory(AbstractAutowireCapableBeanFactory).createBean(String, RootBeanDefinition, Object[]) line: 515	
		DefaultListableBeanFactory(AbstractBeanFactory).lambda$doGetBean$0(String, RootBeanDefinition, Object[]) line: 320	
		603443293.getObject() line: not available	
		DefaultListableBeanFactory(DefaultSingletonBeanRegistry).getSingleton(String, ObjectFactory<?>) line: 222	
		DefaultListableBeanFactory(AbstractBeanFactory).doGetBean(String, Class<T>, Object[], boolean) line: 318	
		DefaultListableBeanFactory(AbstractBeanFactory).getBean(String) line: 199	
		DefaultListableBeanFactory.preInstantiateSingletons() line: 849	
		AnnotationConfigApplicationContext(AbstractApplicationContext).finishBeanFactoryInitialization(ConfigurableListableBeanFactory) line: 877	
		AnnotationConfigApplicationContext(AbstractApplicationContext).refresh() line: 549	
		AnnotationConfigApplicationContext.<init>(Class<?>...) line: 88	
		TestSpringApp.testSelfMyBatis() line: 98	
		TestSpringApp.main(String[]) line: 41	
```

