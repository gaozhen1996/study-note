# 前言

 Class.forName和ClassLoader 都可以来获取Class对象，他们的区别是

- Class.forName 会对类进行初始化（执行静态代码等等）
- ClassLoader不会对类初始化

# 一、效果

Class.forName(String className)这行代码会执行static代码块

```
	//会执行static代码块
	Class.forName(className);
```

ClassLoader不会执行static代码块

```
    //不会执行static代码款
    ClassLoader.getSystemClassLoader().loadClass(className);
```

# 二、分析

## 1.类加载过程

需要先了解类的加载过程

```
装载：通过累的全限定名获取二进制字节流，将二进制字节流转换成方法区中的运行时数据结构，在内存中生成Java.lang.class对象； 
 
链接：执行下面的校验、准备和解析步骤，其中解析步骤是可以选择的； 
 
　　校验：检查导入类或接口的二进制数据的正确性；（文件格式验证，元数据验证，字节码验证，符号引用验证） 
 
　　准备：给类的静态变量分配并初始化存储空间； 
 
　　解析：将常量池中的符号引用转成直接引用； 
 
初始化：激活类的静态变量的初始化Java代码和静态Java代码块，并初始化程序员设置的变量值。
```

## 2.源码分析

Class.forName(className) 调用的底层源码是

```java
    @CallerSensitive
    public static Class<?> forName(String className)
                throws ClassNotFoundException {
        Class<?> caller = Reflection.getCallerClass();
        return forName(className, true, ClassLoader.getClassLoader(caller), caller);
    }

    /**
      * 参数一：className，需要加载的类的名称。
      *	参数二：true，是否对class进行初始化（需要initialize）
      * 参数三：classLoader，对应的类加载器
      */  
    public static Class<?> forName(String name, boolean initialize,
                                       ClassLoader loader)
```

在调用forName方法是，第二个参数选择的true，因此会对类进行初始化，会执行类的静态代码块和静态变量赋值。



完整源码见： https://github.com/gaozhen1996/study-java/blob/master/src/main/java/com/gz/javastudy/javase/reflect/ClassforNameORClassLoader.java 

