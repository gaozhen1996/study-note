- 初始化AnnotationConfigApplicationContext

  - 初始化BeanDefinitionReader

    - ConfigurationClassPostProcessor 重要
    - AutowiredAnnotationBeanPostProcessor 
    - ......

  - 注册配置类

  - refresh()

    - invokeBeanFactoryPostProcessors

      - 执行bean工厂的后置处理器

      - ConfigurationClassPostProcessor 扫描了添加了注解的类

        注册到BeanDefinitionMap中去

    - finishBeanFactoryInitialization

      - getBean的时候检查是不是懒加载的