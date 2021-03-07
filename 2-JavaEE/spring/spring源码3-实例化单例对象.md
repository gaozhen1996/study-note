# 前言

spring实例化对象主要是通过后置处理器来实例化对象的



# 第一次执行BeanPostProcessor

1. **执行方法堆栈：**

- createBean 
  -  resolveBeforeInstantiation 
    -  applyBeanPostProcessorsBeforeInstantiation

2. **处理器接口：**

   InstantiationAwareBeanPostProcessor

3. **执行方法：**

   postProcessBeforeInstantiation

4. **作用：**

   



# 第二次执行BeanPostProcessor

1. **执行方法堆栈**：

- createBean 
  - doCreateBean 
    - createBeanInstance 
      - determineConstructorsFromBeanPostProcessors

2. **处理器接口**

   SmartInstantiationAwareBeanPostProcessor

3. **执行方法**

   determineCandidateConstructors

4. **作用**

   推断实例化bean的构造方法

   

# 第三次执行BeanPostProcessor

1. **执行方法堆栈**：

- createBean 
  - doCreateBean 
    - applyMergedBeanDefinitionPostProcessors 

2. **处理器接口**

   MergedBeanDefinitionPostProcessor

3. **执行方法**

   postProcessMergedBeanDefinition

4. **作用**

   缓存注解信息

# 第四次执行BeanPostProcessor

1. **执行方法堆栈**：

- createBean 
  - doCreateBean 
    - getEarlyBeanReference 

2. **处理器接口**

   SmartInstantiationAwareBeanPostProcessor

3. **执行方法**

   getEarlyBeanReference

4. **作用**

循环引用的后置处理器，获得提前暴露的bean引用。主要用于解决循环引用的问题，只有单例对象才会调用此方法

# 第五次执行BeanPostProcessor

1. **执行方法堆栈**：

- createBean 
  - doCreateBean 
    - populateBean

2. **处理器接口**

   InstantiationAwareBeanPostProcessor

3. **执行方法**

   postProcessAfterInstantiation

4. **作用**

​      该方法的返回值是boolean类型，当返回值为true时，spring容器会自动注入属性值；当返回值为false时，spring容器则不会自动注入属性值。

# 第六次执行BeanPostProcessor

1. **执行方法堆栈**：

- createBean 
  - doCreateBean 
    - populateBean

2. **处理器接口**

   InstantiationAwareBeanPostProcessor

3. **执行方法**

   postProcessProperties

4. **作用**

   完成属性填充

# 第七次执行BeanPostProcessor

1. **执行方法堆栈**：

- createBean 
  - doCreateBean 
    - initializeBean
      - applyBeanPostProcessorsBeforeInitialization

2. **处理器接口**

   BeanPostProcessor

3. **执行方法**

   postProcessBeforeInitialization

4. **作用**



# 第八次执行BeanPostProcessor

1. **执行方法堆栈**：

- createBean 
  - doCreateBean 
    - initializeBean
      - applyBeanPostProcessorsAfterInitialization

2. **处理器接口**

   BeanPostProcessor

3. **执行方法**

   postProcessAfterInitialization

4. **作用**