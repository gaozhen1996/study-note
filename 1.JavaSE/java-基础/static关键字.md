[toc]

# <font color=#2c859d>1．static 修饰变量</font>

<font size=4 color=#999>变量变成类的所属，而不是对象的所属</font>    
    

# <font color=#2c859d>2．static修饰方法</font>

<font size=4 color=#999>方法变成类的所属，而不是对象的所属，可以直接通过类调用方法</font> 


# <font color=#2c859d>3．static修饰静态代码块</font>
<font size=4 color=#999>先执行静态代码快，再执行构造函数,静态代码块只执行一遍　</font> 

```

示例代码：

class Book{
    public Book(String msg) {
        System.out.println(msg);
    }
}

public class Person {
    Book book1 = new Book("book1成员变量初始化");
    static Book book2 ;
    
    public Person(String msg) {
        System.out.println(msg);
    }
    
    Book book3 = new Book("book3成员变量初始化");
    static Book book4 ;
    
    static {
        book2 = new Book("static成员book2成员变量初始化---");
        book4 = new Book("static成员book4成员变量初始化---");
    }
    
    public static void funStatic() {
        System.out.println("static修饰的funStatic方法");
    }
    
    public static void main(String[] args) {
        Person.funStatic();
        System.out.println("****************");
        Person p1 = new Person("p1初始化");
    }
    
}

打印结果：
static成员book2成员变量初始化---
static成员book4成员变量初始化---
static修饰的funStatic方法
****************
book1成员变量初始化
book3成员变量初始化
p1初始化


```

# <font color=#2c859d>４．static修饰静态导包</font>
<font size=4 color=#999>先执行静态代码快，再执行构造函数,静态代码块只执行一遍　</font> 



**PrintHelper.java**
```
package com.gz.testStatic;

public class PrintHelper {
    public static void print(Object o){
        System.out.println(o);
    }
}

```

**App.java**
```
package com.gz.testStatic;

import static com.gz.testStatic.PrintHelper.*;

public class App 
{
    public static void main( String[] args )
    {
        print("Hello World!");
    }
 
}

```

# <font color=#2c859d>5．static修饰类/font>
<font size=4 color=#999>java允许我们在一个类里面定义静态类。比如内部类（nested class）。把nested class封闭起来的类叫外部类。在java中，我们不能用static修饰顶级类（top level class）。只有内部类可以为static。
　</font> 
　
　
```
/* 下面程序演示如何在java中创建静态内部类和非静态内部类 */
class OuterClass{
  private static String msg = "GeeksForGeeks";
  // 静态内部类
  public static class NestedStaticClass{
    // 静态内部类只能访问外部类的静态成员
    public void printMessage() {
     // 试着将msg改成非静态的，这将导致编译错误 
     System.out.println("Message from nested static class: " + msg); 
    }
  }
  // 非静态内部类
  public class InnerClass{
    // 不管是静态方法还是非静态方法都可以在非静态内部类中访问
    public void display(){
     System.out.println("Message from non-static nested class: "+ msg);
    }
  }
} 


class Main
{
  // 怎么创建静态内部类和非静态内部类的实例
  public static void main(String args[]){
    // 创建静态内部类的实例
    OuterClass.NestedStaticClass printer = new OuterClass.NestedStaticClass();
    // 创建静态内部类的非静态方法
    printer.printMessage();  
    // 为了创建非静态内部类，我们需要外部类的实例
    OuterClass outer = new OuterClass();    
    OuterClass.InnerClass inner = outer.new InnerClass();
    // 调用非静态内部类的非静态方法
    inner.display();
    // 我们也可以结合以上步骤，一步创建的内部类实例
    OuterClass.InnerClass innerObject = new OuterClass().new InnerClass();
    // 同样我们现在可以调用内部类方法
    innerObject.display();
  }
}
```