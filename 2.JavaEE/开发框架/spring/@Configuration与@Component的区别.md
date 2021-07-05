# 前言

@Configuration与@Component都可以作为配置类，在使用一般区别不大。但是，在一些特定的状态下，却别还是很大的。

# 一、现象

下面将演示两个配置类，一个是加@Configuration，一个加@Component，在使用情况下的不同。

ConfigReader.java（配置文件读取器类）

```java
import java.io.IOException;
import java.io.InputStream;
import java.util.Properties;

public class ConfigReader {
	private static final String LOCATION = "application.properties";
	
	// 配置信息
	private Properties properties = new Properties();
	
	public ConfigReader() {
		InputStream in = null;
		try {
			in = this.getClass().getClassLoader().getResourceAsStream(LOCATION);
			properties.load(in);
		} catch (IOException e) {
			e.printStackTrace();
		} finally {
			try {
				if (in != null)
					in.close();
			} catch (IOException e) {
				e.printStackTrace();
			}
		}
	}
	
	public String getconfigValueByKey(String key) {
		return properties.getProperty(key);
	}
}
```



ConfigurationConfig.java(添加@Configuration的配置类)

```java
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class ConfigurationConfig {
	
	@Bean
	public ConfigReader getConfigReader() {
		return new ConfigReader();
	}
}
```

ComponentConfig.java(添加@Component的配置类)

```java
import org.springframework.context.annotation.Bean;
import org.springframework.stereotype.Component;

@Component
public class ComponentConfig {
	
	@Bean
	public ConfigReader getConfigReader() {
		return new ConfigReader();
	}
}
```

Main.java(main方法)

```java
    public static void main(String[] args) {
    	AnnotationConfigApplicationContext context = new    		                    
                 AnnotationConfigApplicationContext(TestMain.class);
  
        ConfigurationConfig configurationConfig  = 
            (ConfigurationConfig) context.getBean("configurationConfig");
        System.out.println("使用@Configuration");
        System.out.println(configurationConfig.getConfigReader()
                           ==configurationConfig.getConfigReader());
        
        ComponentConfig componentConfig  =
             (ComponentConfig) context.getBean("componentConfig");
        System.out.println("使用@Component");
        System.out.println(componentConfig.getConfigReader()
                           ==componentConfig.getConfigReader());
        
        Scanner scanner = new Scanner(System.in);
        while (true) {
            System.out.println("加载文件请输入1，退出请输入0");
            String key = scanner.nextLine();
            if(key.equals("0")) {
            	scanner.close();
            	context.close();
            	System.exit(0);
            }
            System.out.println(configurationConfig.getConfigReader()
                               .getconfigValueByKey("scanPackage"));
			System.out.println(componentConfig.getConfigReader()
                               .getconfigValueByKey("scanPackage"));
		}
	}
```

输出

> 使用@Configuration
> true
> 使用@Component
> false
> 加载文件请输入1，退出请输入0
> 输入1
> com.gz.javastudy.springsimple.demo
> com.gz.javastudy.springsimple.demo
> 加载文件请输入1，退出请输入0
> 修改配置文件，重新输入1
> com.gz.javastudy.springsimple.demo
> com.gz.javastudy.springsimple.demo1

**结论：可以看出，使用@Configuration返回的是同样的对象，使用@Component返回的对是new的**

因此在修改配置文件后，加@Configuration不能够更新

# 二、原因分析

打印一下两个配置类，可见：

com.gz.javastudy.springapp.config.ConfigurationConfig$$EnhancerBySpringCGLIB$$4bb2c78@780cb77

com.gz.javastudy.springapp.config.ComponentConfig@161b062a

加了@Configuration已经是被CGLIB代理过的类。

## 1.@Configuration生成代理对象的源码分析

- 第一步：在执行postProcessBeanDefinitionRegistry时， 解析配置类的时候,会将加了@Configuration设置full的标志。

- 第二步：在执行postProcessBeanFactory时，会对加了full的配置类，使用cglib生成代理对象。

### 先添加标志

ConfigurationClassPostProcessor.java

```java
public void processConfigBeanDefinitions(BeanDefinitionRegistry registry) {
		List<BeanDefinitionHolder> configCandidates = new ArrayList<>();
		String[] candidateNames = registry.getBeanDefinitionNames();

		for (String beanName : candidateNames) {
			BeanDefinition beanDef = registry.getBeanDefinition(beanName);
			if (ConfigurationClassUtils.isFullConfigurationClass(beanDef) ||
					ConfigurationClassUtils.isLiteConfigurationClass(beanDef)) {
				if (logger.isDebugEnabled()) {
					logger.debug("Bean definition has already been processed as a                                 configuration class: " + beanDef);
				}
			}
            //判断是否是配置类
			else if (ConfigurationClassUtils.checkConfigurationClassCandidate(beanDef,                           this.metadataReaderFactory)) {
				configCandidates.add(new BeanDefinitionHolder(beanDef, beanName));
			}
		}
}
```

ConfigurationClassUtils.java

```java
public static boolean checkConfigurationClassCandidate(
		BeanDefinition beanDef, MetadataReaderFactory metadataReaderFactory) {
  //如果是加了@Configuration注解的，则在beanDefinition中设置一个full的标志     
  if (isFullConfigurationCandidate(metadata)) {
	 beanDef.setAttribute(CONFIGURATION_CLASS_ATTRIBUTE, 		                   	                        CONFIGURATION_CLASS_FULL);
   }
   //如果是加了@Configuration注解的，则在beanDefinition中设置一个lite的标志
   else if (isLiteConfigurationCandidate(metadata)) {
			beanDef.setAttribute(CONFIGURATION_CLASS_ATTRIBUTE,                                                                CONFIGURATION_CLASS_LITE);
	}
	else {
		return false;
	}
}
```

### 判断是否需要增强

ConfigurationClassPostProcessor.java

```java
public void postProcessBeanFactory(ConfigurableListableBeanFactory beanFactory) {
    //忽略其他代码
    enhanceConfigurationClasses(beanFactory);
}

public void enhanceConfigurationClasses(ConfigurableListableBeanFactory beanFactory) {
    Map<String, AbstractBeanDefinition> configBeanDefs = new LinkedHashMap<>();
    for (String beanName : beanFactory.getBeanDefinitionNames()) {
        BeanDefinition beanDef = beanFactory.getBeanDefinition(beanName);
        if (ConfigurationClassUtils.isFullConfigurationClass(beanDef)) {
            //如果是full配置类，则先添加到configBeanDefs中
            configBeanDefs.put(beanName, (AbstractBeanDefinition) beanDef);
        }
    }
    if (configBeanDefs.isEmpty()) {
        // nothing to enhance -> return immediately
        return;
    }

    ConfigurationClassEnhancer enhancer = new ConfigurationClassEnhancer();
    for (Map.Entry<String, AbstractBeanDefinition> entry : configBeanDefs.entrySet()){   
        AbstractBeanDefinition beanDef = entry.getValue();
    	//设置为cglib代理对象
        beanDef.setBeanClass(enhancedClass);
    }
}
```

反编译代理后的代码，getConfigReader()代码已经不是原本写的代码了。

```java
   public final ConfigReader getConfigReader() {
      MethodInterceptor var10000 = this.CGLIB$CALLBACK_0;
      if (this.CGLIB$CALLBACK_0 == null) {
         CGLIB$BIND_CALLBACKS(this);
         var10000 = this.CGLIB$CALLBACK_0;
      }

      return var10000 != null ? (ConfigReader)var10000.intercept(this, 		                               CGLIB$getConfigReader$0$Method
                  , CGLIB$emptyArgs, CGLIB$getConfigReader$0$Proxy) :            		                   super.getConfigReader();
   }
```

