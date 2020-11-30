## 一、编程规约

### 命名风格

1. 【强制】代码中的命名均不能以下划线或美元符号开始，也不能以下划线或美元符号结束。

   > 反例：_name/ __name / $name / name_ / name$ / name__

2. 【强制】代码中的命名严禁使用拼音与英文混合的方式，更不允许直接使用中文的方式。

   > 说明：正确的英文拼写和语法可以让阅读者易于理解，避免歧义。注意，即使纯拼音命名方式也要避免采用。
   > 正例：hundsun / hangzhou 等国际通用的名称，可视同英文。
   > 反例：DaZhePromotion [打折] / getPingfenByName() [评分] / int 某变量 = 3

3. 【强制】类名使用 UpperCamelCase 风格，但以下情形例外：DO / BO / DTO / VO / AO / PO / UID 等。

   > 正例：MarcoPolo / UserDO / XmlService / TcpUdpDeal / TaPromotion
   > 反例：macroPolo / UserDo / XMLService / TCPUDPDeal / TAPromotion

4. 【强制】方法名、参数名、成员变量、局部变量都统一使用 lowerCamelCase风格，必须遵从驼峰形式。

   > 正例： localValue / getHttpMessage() / inputUserId

5. 【强制】常量命名全部大写，单词间用下划线隔开，力求语义表达完整清楚，不要嫌名字长。

   > 正例：MAX_STOCK_COUNT
   > 反例：MAX_COUNT

6. 【强制】类型与中括号紧挨相连来表示数组。

   > 正例：定义整形数组 int[] arrayDemo;
   > 反例：在 main 参数中，使用 String args[]来定义。

7. 【强制】POJO 类中布尔类型的变量，都不要加 is前缀，否则部分框架解析会引起序列化错误。

   > 反例：定义为基本数据类型 Boolean isDeleted 的属性，它的方法也是isDeleted()，RPC框架在反向解析的时候，“误以为”对应的属性名称是deleted，导致属性获取不到，进而抛出异常。

8. 【强制】包名统一使用小写，点分隔符之间有且仅有一个自然语义的英语单词。包名统一使用单数形式，但是类名如果有复数含义，类名可以使用复数形式。

   > 正例：应用工具类包名为 com.alibaba.ai.util、类名为 MessageUtils（此规则参考spring的框架结构）

9. 【强制】杜绝完全不规范的缩写，避免望文不知义。

   > 反例：AbstractClass“缩写”命名成 AbsClass；condition“缩写”命名成condi，此类随意缩写严重降低了代码的可阅读性。

10. 接口和实现类的命名有两套规则：

    > 1）【强制】对于 Service 和 DAO 类，基于 SOA的理念，暴露出来的服务一定是接口，内部的实现类用 Impl 的后缀与接口区别。
    > 正例：CacheServiceImpl 实现 CacheService 接口。
    > 2）【推荐】 如果是形容能力的接口名称，取对应的形容词为接口名（通常是–able的形式）。
    > 正例：AbstractTranslator 实现 Translatable 接口。

11. 【参考】枚举类名建议带上 Enum后缀，枚举成员名称需要全大写，单词间用下划线隔开。

    > 说明：枚举其实就是特殊的类，域成员均为常量，且构造方法被默认强制是私有。
    > 正例：枚举名字为 ProcessStatusEnum 的成员名称：SUCCESS / UNKNOWN_REASON。

12. 【参考】各层命名规约：

    ```
    1.  Service/DAO 层方法命名规约        
       1） 获取单个对象的方法用 get 做前缀。            
       2） 获取多个对象的方法用 list 做前缀，复数形式结尾如：listObjects。              
       3） 获取统计值的方法用 count 做前缀。           
       4） 插入的方法用 save/insert 做前缀。              
       5） 删除的方法用 remove/delete 做前缀。                        
       6） 修改的方法用 update 做前缀。    
    2.  领域模型命名规约   
       1） 数据对象：xxxDO，xxx 即为数据表名。   
       2） 数据传输对象：xxxDTO，xxx 为业务领域相关的名称。   
       3） 展示对象：xxxVO，xxx 一般为网页名称。   
       4） POJO 是 DO/DTO/BO/VO 的统称，禁止命名成 xxxPOJO    
    ```

### 常量定义

1. 【强制】不允许任何魔法值（即未经预先定义的常量）直接出现在代码中。

   > 反例：String key = "Id#taobao_" + tradeId;cache.put(key, value);

2. 【强制】在 long 或者 Long 赋值时，数值后使用大写的 L，不能是小写的l，小写容易跟数字 1 混淆，造成误解。

   > 说明：Long a = 2l; 写的是数字的 21，还是 Long 型的 2?

3. 【推荐】不要使用一个常量类维护所有常量，要按常量功能进行归类，分开维护。

   > 说明：大而全的常量类，杂乱无章，使用查找功能才能定位到修改的常量，不利于理解和维护。
   > 正例：缓存相关常量放在类 CacheConsts 下；系统配置相关常量放在类 ConfigConsts下。

4. 【推荐】常量的复用层次有五层：跨应用共享常量、应用内共享常量、子工程内共享常量、包内共享常量、类内共享常量。
   1） 跨应用共享常量：放置在二方库中，通常是 client.jar 中的 constant 目录下。
   2） 应用内共享常量：放置在一方库中，通常是子模块中的 constant 目录下。

   > 反例：易懂变量也要统一定义成应用内共享常量，两位攻城师在两个类中分别定义了表示“是”的变量：
   > A 中：public static final String YES = "yes";
   > B 中：public static final String YES = "y";
   > A.YES.equals(B.YES)，预期是 true，但实际返回为 false，导致线上问题。

   3） 子工程内部共享常量：即在当前子工程的 constant 目录下。
   4） 包内共享常量：即在当前包下单独的 constant 目录下。
   5） 类内共享常量：直接在类内部 private static final 定义。

5. 【推荐】如果变量值仅在一个固定范围内变化用 enum 类型来定义。

   > 说明：如果存在名称之外的延伸属性应使用 enum类型，下面正例中的数字就是延伸信息，表示一年中的第几个季节。
   > 正例：
   >
   > ```
   >    public enum SeasonEnum {
   >        SPRING(1), SUMMER(2), AUTUMN(3), WINTER(4);
   >        private int seq;
   >        SeasonEnum(int seq){
   >        this.seq = seq;
   >        }
   >    }
   > ```

### 代码格式

1. 【强制】大括号的使用约定。如果是大括号内为空，则简洁地写成{}即可，不需要换行；如果是非空代码块则：
   1） 左大括号前不换行。
   2） 左大括号后换行。
   3） 右大括号前换行。
   4） 右大括号后还有 else 等代码则不换行；表示终止的右大括号后必须换行。

2. 【强制】左小括号和字符之间不出现空格；同样，右小括号和字符之间也不出现空格；而左大括号前需要空格。详见第5条下方

   正例

   提示。

   > 反例：if (空格a == b空格)

3. 【强制】if/for/while/switch/do 等保留字与括号之间都必须加空格。

4. 【强制】任何二目、三目运算符的左右两边都需要加一个空格。

   > 说明：运算符包括赋值运算符=、逻辑运算符&&、加减乘除符号等。

5. 【强制】采用 4 个空格缩进，禁止使用 tab 字符。

   > 说明：如果使用 tab 缩进，必须设置 1 个 tab 为 4 个空格。IDEA 设置 tab 为4个空格时，请勿勾选 Use tab character；而在 eclipse 中，必须勾选 insert spaces for tabs。
   > 正例： （涉及 1-5 点）
   >
   > ```
   > public static void main(String[] args) {
   >     // 缩进 4 个空格
   >     String say = "hello";
   >     // 运算符的左右必须有一个空格 
   >     int flag = 0;
   >     // 关键词 if 与括号之间必须有一个空格，括号内的 f 与左括号，0 与右括号不需要空格 
   >     if (flag == 0) {
   >         System.out.println(say);
   >     }
   >     // 左大括号前加空格且不换行；左大括号后换行
   >     if (flag == 1) {
   >         System.out.println("world");
   >     // 右大括号前换行，右大括号后有 else，不用换行
   >     } else { 
   >         System.out.println("ok");
   >     // 在右大括号后直接结束，则必须换行
   >     }
   > }
   > ```

6. 【强制】方法参数在定义和传入时，多个参数逗号后边必须加空格。

   > 正例：下例中实参的 args1，后边必须要有一个空格。
   > method(args1, args2, args3);

7. 【强制】IDE 的 text file encoding 设置为 UTF-8; IDE 中文件的换行符使用 Unix 格式，不要使用 Windows 格式。

8. 【推荐】单个方法的总行数不超过 80 行。

   > 说明：包括方法签名、结束右大括号、方法内代码、注释、空行、回车及任何不可见字符的总行数不超过 80 行。
   > 正例：代码逻辑分清红花和绿叶，个性和共性，绿叶逻辑单独出来成为额外方法，使主干代码更加清晰；共性逻辑抽取成为共性方法，便于复用和维护。

9. 【推荐】没有必要增加若干空格来使某一行的字符与上一行对应位置的字符对齐。

   > 正例：
   >
   > ```
   >    int one = 1;
   >    long two = 2L;
   >    float three = 3F;
   >    StringBuffer sb = new StringBuffer();
   > ```
   >
   > 说明：增加 sb 这个变量，如果需要对齐，则给 a、b、c都要增加几个空格，在变量比较多的情况下，是非常累赘的事情。

10. 【推荐】不同逻辑、不同语义、不同业务的代码之间插入一个空行分隔开来以提升可读性。

    > 说明：任何情形，没有必要插入多个空行进行隔开。

11. 【推荐】注释的双斜线与注释内容之间有且仅有一个空格。

    > 正例：
    >
    > ```
    >    // 这是示例注释，请注意在双斜线之后有一个空格
    >    String ygb = new String();
    > ```

12. 【推荐】单行字符数限制不超过 120 个，超出需要换行，换行时遵循如下原则：
    1） 第二行相对第一行缩进 4 个空格，从第三行开始，不再继续缩进，参考示例。
    2） 运算符与下文一起换行。
    3） 方法调用的点符号与下文一起换行。
    4） 方法调用中的多个参数需要换行时，在逗号后进行。
    5） 在括号前不要换行，见反例。

    > 正例：

    ```
          StringBuffer sb = new StringBuffer();
    
          // 超过 120 个字符的情况下，换行缩进 4 个空格，点号和方法名称一起换行
          sb.append("zi").append("xin")...
              .append("huang")...
              .append("huang")...
              .append("huang");
    ```

    > 反例：

    ```
          StringBuffer sb = new StringBuffer();
    
          // 超过 120 个字符的情况下，不要在括号前换行
          sb.append("zi").append("xin")...append
          ("huang");
    
          // 参数很多的方法调用可能超过 120 个字符，不要在逗号前换行 
          method(args1, args2, args3, ...
               , argsX);
    ```

### OOP 规约

1. 【强制】避免通过一个类的对象引用访问此类的静态变量或静态方法，无谓增加编译器解析成本，直接用类名来访问即可。

2. 【强制】所有的覆写方法，必须加@Override 注解。

   > 说明：getObject()与 get0bject()的问题。一个是字母的 O，一个是数字的 0，加 @Override 可以准确判断是否覆盖成功。另外，如果在抽象类中对方法签名进行修改，其实现类会马上编译报错。

3. 【强制】相同参数类型，相同业务含义，才可以使用 Java 的可变参数，避免使用 Object。

   > 说明：可变参数必须放置在参数列表的最后。（提倡同学们尽量不用可变参数编程）
   > 正例：public List listUsers(String type, Long... ids) {...}

4. 【强制】外部正在调用或者二方库依赖的接口，不允许修改方法签名，避免对接口调用方产生影响。接口过时必须加@Deprecated注解，并清晰地说明采用的新接口或者新服务是什么。

5. 【强制】不能使用过时的类或方法。

   > 说明：java.net.URLDecoder 中的方法 decode(String encodeStr) 这个方法已经过时，应该使用双参数 decode(String source, String encode)。接口提供方既然明确是过时接口， 那么有义务同时提供新的接口；作为调用方来说，有义务去考证过时方法的新实现是什么。

6. 【强制】Object 的 equals 方法容易抛空指针异常，应使用常量或确定有值的对象来调用 equals。

   > 正例："test".equals(object);
   > 反例：object.equals("test");
   > 说明：推荐使用 java.util.Objects#equals（JDK7 引入的工具类）

7. 【强制】所有的相同类型的包装类对象之间值的比较，全部使用 equals 方法比较。

   > 说明：对于 Integer var = ? 在-128 至 127 范围内的赋值，Integer 对象是在 IntegerCache.cache 产生，会复用已有对象，这个区间内的 Integer值可以直接使用==进行判断，但是这个区间之外的所有数据，都会在堆上产生，并不会复用已有对象，这是一个大坑，推荐使用equals 方法进行判断。

8. 关于基本数据类型与包装数据类型的使用标准如下：

   1） 【强制】所有的 POJO 类属性必须使用包装数据类型。

   2） 【强制】RPC 方法的返回值和参数必须使用包装数据类型。

   3） 【推荐】所有的局部变量使用基本数据类型。

   > 说明：POJO类属性没有初值是提醒使用者在需要使用时，必须自己显式地进行赋值，任何NPE 问题，或者入库检查，都由使用者来保证。
   > 正例：数据库的查询结果可能是 null，因为自动拆箱，用基本数据类型接收有 NPE 风险。
   > 反例：比如显示成交总额涨跌情况，即正负 x%，x 为基本数据类型，调用的 RPC 服务，调用不成功时，返回的是默认值，页面显示为0%，这是不合理的，应该显示成中划线。所以包装数据类型的 null 值，能够表示额外的信息，如：远程调用失败，异常退出。

9. 【推荐】定义 DO/DTO/VO 等 POJO 类时，不要设定任何属性默认值。

   > 反例：POJO 类的 gmtCreate 默认值为 new Date()，但是这个属性在数据提取时并没有置入具体值，在更新其它字段时又附带更新了此字段，导致创建时间被修改成当前时间。

10. 【强制】POJO 类必须写 toString 方法。使用 IDE 中的工具：source> generate toString 时，如果继承了另一个 POJO 类，注意在前面加一下 super.toString。

    > 说明：在方法执行抛出异常时，可以直接调用 POJO 的 toString()方法打印其属性值，便于排查问题。

11. 【强制】禁止在 POJO 类中，同时存在对应属性 xxx 的 isXxx()和 getXxx()方法。

    > 说明：框架在调用属性 xxx 的提取方法时，并不能确定哪个方法一定是被优先调用到。

12. 【推荐】当一个类有多个构造方法，或者多个同名方法，这些方法应该按顺序放置在一起，便于阅读，此条规则优先于第16 条规则。

13. 【推荐】 类内方法定义的顺序依次是：公有方法或保护方法 > 私有方法 > getter/setter方法。

    > 说明：公有方法是类的调用者和维护者最关心的方法，首屏展示最好；保护方法虽然只是子类关心，也可能是“模板设计模式”下的核心方法；而私有方法外部一般不需要特别关心，是一个黑盒实现；因为承载的信息价值较低，所有 Service 和 DAO 的 getter/setter 方法放在类体最后。

14. 【推荐】setter 方法中，参数名称与类成员变量名称一致，this.成员名 = 参数名。在 getter/setter 方法中，不要增加业务逻辑，增加排查问题的难度。

    > 反例：
    >
    > ```
    >       public Integer getData() {
    >           if (condition) {
    >               return this.data + 100;
    >           } else {
    >               return this.data - 100;
    >           }
    >       }
    > ```

15. 【推荐】循环体内，字符串的连接方式，使用 StringBuilder 的 append 方法进行扩展。

    > 说明：下例中，反编译出的字节码文件显示每次循环都会 new 出一个 StringBuilder 对象，然后进行 append 操作，最后通过 toString 方法返回 String 对象，造成内存资源浪费。
    > 反例：
    >
    > ```
    >       String str = "start";
    >       for (int i = 0; i < 100; i++) {
    >            str = str + "hello";
    >       }
    > ```

### 集合处理

1. 【强制】关于 hashCode 和 equals 的处理，遵循如下规则：

   1） 只要重写 equals，就必须重写 hashCode。

   2） 因为 Set 存储的是不重复的对象，依据 hashCode 和 equals 进行判断，所以 Set 存储的对象必须重写这两个方法。

   3） 如果自定义对象作为 Map 的键，那么必须重写 hashCode 和 equals。

   > 说明：String 重写了 hashCode 和 equals 方法，所以我们可以非常愉快地使用String 对象作为 key 来使用。

2. 【强制】ArrayList 的 subList 结果不可强转成 ArrayList，否则会抛出 ClassCastException 异常，即 java.util.RandomAccessSubList cannot be cast to java.util.ArrayList。

   > 说明：subList 返回的是 ArrayList 的内部类 SubList，并不是 ArrayList 而是 ArrayList 的一个视图，对于 SubList 子列表的所有操作最终会反映到原列表上。

3. 【强制】在 subList 场景中，高度注意对原集合元素的增加或删除，均会导致子列表的遍历、增加、删除产生 ConcurrentModificationException 异常。

4. 【强制】使用集合转数组的方法，必须使用集合的 toArray(T[] array)，传入的是类型完全一样的数组，大小就是 list.size()。

   > 说明：使用 toArray 带参方法，入参分配的数组空间不够大时，toArray 方法内部将重新分配内存空间，并返回新数组地址；如果数组元素个数大于实际所需，下标为[ list.size() ] 的数组元素将被置为 null，其它数组元素保持原值，因此最好将方法入参数组大小定义与集合元素个数一致。
   > 正例：
   >
   > ```
   >    List<String> list = new ArrayList<String>(2);
   >    list.add("guan");
   >    list.add("bao");
   >    String[] array = new String[list.size()];
   >    array = list.toArray(array);
   > ```
   >
   > 反例：直接使用 toArray 无参方法存在问题，此方法返回值只能是Object[]类，若强转其它类型数组将出现 ClassCastException 错误。

5. 【强制】使用工具类Arrays.asList()把数组转换成集合时，不能使用其修改集合相关的方法，它的 add/remove/clear 方法会抛出 UnsupportedOperationException 异常。

   > 说明：asList 的返回对象是一个 Arrays内部类，并没有实现集合的修改方法。Arrays.asList体现的是适配器模式，只是转换接口，后台的数据仍是数组。 String[] str = new String[] { "you", "wu" }; List list = Arrays.asList(str); 第一种情况：list.add("yangguanbao"); 运行时异常。 第二种情况：str[0] = "gujin"; 那么 list.get(0)也会随之修改。

6. 【强制】泛型通配符<? extends T>来接收返回的数据，此写法的泛型集合不能使用 add 方法，而<? super T>不能使用 get 方法，作为接口调用赋值时易出错。

   > 说明：扩展说一下 PECS(Producer Extends Consumer Super)原则：第一、频繁往外读取内容的，适合用<? extends T>。第二、经常往里插入的，适合用<? super T>。

7. 【强制】不要在 foreach 循环里进行元素的 remove/add 操作。remove 元素请使用 Iterator方式，如果并发操作，需要对 Iterator 对象加锁。

   > 正例：
   >
   > ```
   >     List<String> list = new ArrayList<>();
   >     list.add("1");
   >     list.add("2");
   >     Iterator<String> iterator = list.iterator();
   >     while (iterator.hasNext()) {
   >         String item = iterator.next();
   >         if (删除元素的条件) {
   >             iterator.remove();
   >         }
   >     }
   > ```
   >
   > 反例：
   >
   > ```
   >    for (String item : list) {
   >        if ("1".equals(item)) {
   >             list.remove(item);
   >        }
   >    }
   > ```
   >
   > 说明：以上代码的执行结果肯定会出乎大家的意料，那么试一下把“1”换成“2”，会是同样的结果吗？

8. 【强制】 在 JDK7 版本及以上，Comparator 实现类要满足如下三个条件，不然 Arrays.sort，Collections.sort 会报 IllegalArgumentException 异常。

   > 说明：三个条件如下
   > 1） x，y 的比较结果和 y，x 的比较结果相反。
   > 2） x>y，y>z，则 x>z。
   > 3） x=y，则 x，z 比较结果和 y，z 比较结果相同。
   > 反例：下例中没有处理相等的情况，实际使用中可能会出现异常：
   >
   > ```
   >    new Comparator<Student>() {
   >        @Override
   >        public int compare(Student o1, Student o2) {
   >            return o1.getId() > o2.getId() ? 1 : -1;
   >        }
   >    };
   > ```

9. 【推荐】集合泛型定义时，在 JDK7 及以上，使用 diamond 语法或全省略。

   > 说明：菱形泛型，即 diamond，直接使用\<>来指代前边已经指定的类型。
   > 正例：
   > // <> diamond 方式
   > HashMap<String, String> userCache = new HashMap<>(16);
   > // 全省略方式
   > ArrayList<User> users = new ArrayList(10);

10. 【推荐】集合初始化时，指定集合初始值大小。

    > 说明：HashMap 使用 HashMap(int initialCapacity) 初始化。
    > 正例：initialCapacity = (需要存储的元素个数 / 负载因子) + 1。注意负载因子（即 loader factor）默认为 0.75，如果暂时无法确定初始值大小，请设置为 16（即默认值）。
    > 反例：HashMap 需要放置 1024 个元素，由于没有设置容量初始大小，随着元素不断增加，容 7 次被迫扩大，resize 需要重建 hash 表，严重影响性能。

11. 【推荐】使用 entrySet 遍历 Map 类集合 KV，而不是 keySet 方式进行遍历。

    > 说明：keySet 其实是遍历了 2 次，一次是转为 Iterator 对象，另一次是从 hashMap 中取出 key 所对应的 value。而 entrySet 只是遍历了一次就把 key 和 value 都放到了 entry 中，效 率更高。如果是 JDK8，使用 Map.foreach 方法。
    > 正例：values()返回的是 V 值集合，是一个 list 集合对象；keySet()返回的是 K 值集合，是一个 Set 集合对象；entrySet()返回的是 K-V 值组合集合。

12. 【推荐】高度注意 Map 类集合 K/V 能不能存储 null 值的情况，如下表格：

| **集合类**        | **Key**       | **Value**     | **Super**   | 说明                   |
| :---------------- | :------------ | :------------ | :---------- | :--------------------- |
| Hashtable         | 不允许为 null | 不允许为 null | Dictionary  | 线程安全               |
| ConcurrentHashMap | 不允许为 null | 不允许为 null | AbstractMap | 锁分段技术（JDK8:CAS） |
| TreeMap           | 不允许为 null | 允许为 null   | AbstractMap | 线程不安全             |
| HashMap           | 允许为 null   | 允许为 null   | AbstractMap | 线程不安全             |

> 反例： 由于 HashMap 的干扰，很多人认为 ConcurrentHashMap 是可以置入 null 值，而事实上， 存储 null 值时会抛出 NPE 异常。

1. 【参考】合理利用好集合的有序性(sort)和稳定性(order)，避免集合的无序性(unsort)和不稳定性(unorder)带来的负面影响。

   > 说明：有序性是指遍历的结果是按某种比较规则依次排列的。稳定性指集合每次遍历的元素次序是一定的。如：ArrayList 是 order/unsort；HashMap 是 unorder/unsort；TreeSet 是order/sort。

### 并发处理

1. 【强制】创建线程或线程池时请指定有意义的线程名称，方便出错时回溯。

   > 正例：
   >
   > ```
   >    public class TimerTaskThread extends Thread {
   >        public TimerTaskThread() {
   >            super.setName("TimerTaskThread");
   >            ...
   >        }
   >    }
   > ```

2. 【强制】线程资源必须通过线程池提供，不允许在应用中自行显式创建线程。

   > 说明：使用线程池的好处是减少在创建和销毁线程上所消耗的时间以及系统资源的开销，解决资源不足的问题。如果不使用线程池，有可能造成系统创建大量同类线程而导致消耗完内存或者“过度切换”的问题。

3. 【强制】线程池不允许使用 Executors 去创建，而是通过 ThreadPoolExecutor 的方式，这样的处理方式让写的同学更加明确线程池的运行规则，规避资源耗尽的风险。

   > 说明：Executors 返回的线程池对象的弊端如下：
   > 1）FixedThreadPool 和 SingleThreadPool:
   > 允许的请求队列长度为 Integer.MAX_VALUE，可能会堆积大量的请求，从而导致 OOM。
   > 2）CachedThreadPool 和 ScheduledThreadPool:
   > 允许的创建线程数量为 Integer.MAX_VALUE，可能会创建大量的线程，从而导致 OOM。

4. 【强制】使用 Instant 代替 Date，LocalDateTime 代替Calendar，DateTimeFormatter 代替SimpleDateFormat。

   > 说明：SimpleDateFormat是线程不安全的类，在JDK8的应用中，替代类有使用简单、代码优雅、线程安全优点

5. 【强制】对多个资源、数据库表、对象同时加锁时，需要保持一致的加锁顺序，否则可能会造成死锁。

   > 说明：线程一需要对表 A、B、C 依次全部加锁后才可以进行更新操作，那么线程二的加锁顺序也必须是 A、B、C，否则可能出现死锁。

6. 【推荐】每次访问冲突概率小于 20%，推荐使用乐观锁，否则使用悲观锁。乐观锁的重试次数不得小于 3 次。

   > 说明：乐观锁顾名思义就是很乐观，每次去拿数据的时候都认为别人不会修改，所以不会上锁，但是在更新的时候会判断一下在此期间别人有没有去更新这个数据，冲突频率高的话不应选择乐观锁，需要重试好几次，代价大。

7. 【强制】多线程并行处理定时任务时，Timer 运行多个 TimeTask 时，只要其中之一没有捕获抛出的异常，其它任务便会自动终止运行，使用 ScheduledExecutorService 则没有这个问题。

8. 【推荐】使用 CountDownLatch 进行异步转同步操作，每个线程退出前必须调用 countDown 方法，线程执行代码注意 catch 异常，确保 countDown 方法被执行到，避免主线程无法执行至 await 方法，直到超时才返回结果。

   > 说明：注意，子线程抛出异常堆栈，不能在主线程 try-catch 到。

9. 【推荐】避免 Random实例被多线程使用，虽然共享该实例是线程安全的，但会因竞争同一 seed 导致的性能下降。

   > 说明：Random 实例包括 java.util.Random 的实例或者 Math.random()的方式。
   > 正例：在 JDK7 之后，可以直接使用 API ThreadLocalRandom，而在 JDK7 之前，需要编码保证每个线程持有一个实例。

10. 【推荐】在并发场景下，通过双重检查锁（double-checked locking）实现延迟初始化的优化问题隐患(可参考 The "Double-Checked Locking is Broken" Declaration)，推荐解决方案中较为简单一种（适用于 JDK5 及以上版本），将目标属性声明为 volatile 型。

    > 反例：
    >
    > ```
    >       class LazyInitDemo {
    >           private Helper helper = null;
    >           public Helper getHelper() {
    >               if (helper == null) synchronized(this) {
    >                   if (helper == null)
    >                   helper = new Helper();
    >               }
    >               return helper;
    >           }
    >       // other methods and fields...
    >       }
    > ```

11. 【参考】volatile 解决多线程内存不可见问题。对于一写多读，是可以解决变量同步问题， 但是如果多写，同样无法解决线程安全问题。如果是 count++操作，使用如下类实现： AtomicInteger count = new AtomicInteger(); count.addAndGet(1); 如果是 JDK8，推荐使用 LongAdder 对象，比 AtomicLong 性能更好（减少乐观锁的重试次数）。

12. 【参考】 HashMap 在容量不够进行 resize 时由于高并发可能出现死链，导致 CPU 飙升，在开发过程中可以使用其它数据结构或加锁来规避此风险。

13. 【参考】ThreadLocal 无法解决共享对象的更新问题，ThreadLocal 对象建议使用 static修饰。这个变量是针对一个线程内所有操作共享的，所以设置为静态变量，所有此类实例共享 此静态变量，也就是说在类第一次被使用时装载，只分配一块存储空间，所有此类的对象(只要是这个线程内定义的)都可以操控这个变量。

### 控制语句

1. 【强制】在一个 switch 块内，每个 case 要么通过 break/return 等来终止，要么注释说明程序将继续执行到哪一个 case 为止；在一个 switch 块内，都必须包含一个 default 语句并且放在最后，即使空代码。

2. 【强制】在 if/else/for/while/do 语句中必须使用大括号。即使只有一行代码，避免采用单行的编码方式：if (condition) statements;

3. 【强制】超过 3 层的 if-else 的逻辑判断代码可以使用卫语句、策略模式、状态模式等来实现

   > 说明：避免后续代码维护困难，请勿超过 3 层。
   > 正例：卫语句示例如下：
   >
   > ```
   >    public void today() {
   >        if (isBusy()) {
   >            System.out.println(“change time.”);
   >            return;
   >        }
   >        if (isFree()) {
   >            System.out.println(“go to travel.”);
   >            return;
   >        }
   >        System.out.println(“stay at home to learn HUNDSUN Java Coding Guidelines.”);
   >        return;
   >    }
   > ```

4. 【强制】循环体中的语句要考量性能，以下操作尽量移至循环体外处理，如定义对象、变量、获取数据库连接，进行不必要的 try-catch 操作（这个 try-catch 是否可以移至循环体外）。

5. 【推荐】避免采用取反逻辑运算符。

   > 说明：取反逻辑不利于快速理解，并且取反逻辑写法必然存在对应的正向逻辑写法。
   > 正例：使用 if (x < 628) 来表达 x 小于 628。
   > 反例：使用 if (!(x >= 628)) 来表达 x 小于 628。

6. 【推荐】接口入参保护，这种场景常见的是用作批量操作的接口。

7. 【参考】下列情形，需要进行参数校验：
   1） 调用频次低的方法。
   2） 执行时间开销很大的方法。此情形中，参数校验时间几乎可以忽略不计，但如果因为参 数错误导致中间执行回退，或者错误，那得不偿失。
   3） 需要极高稳定性和可用性的方法。
   4） 对外提供的开放接口，不管是 RPC/API/HTTP 接口。
   5） 敏感权限入口。

8. 【参考】下列情形，不需要进行参数校验：
   1） 极有可能被循环调用的方法。但在方法说明里必须注明外部参数检查要求。
   2） 底层调用频度比较高的方法。毕竟是像纯净水过滤的最后一道，参数错误不太可能到底层才会暴露问题。一般 DAO 层与 Service 层都在同一个应用中，部署在同一台服务器中，所以 DAO 的参数校验，可以省略。
   3） 被声明成 private只会被自己代码所调用的方法，如果能够确定调用方法的代码传入参数已经做过检查或者肯定不会有问题，此时可以不校验参数。

### 注释规约

1. 【强制】类、类属性、类方法的注释必须使用 Javadoc 规范，使用/**内容*/格式，不得使用 // xxx 方式。

   > 说明：在 IDE 编辑窗口中，Javadoc 方式会提示相关注释，生成 Javadoc 可以正确输出相应注释；在 IDE 中，工程调用方法时，不进入方法即可悬浮提示方法、参数、返回值的意义，提高阅读效率。

2. 【强制】所有的抽象方法（包括接口中的方法）必须要用 Javadoc 注释、除了返回值、参数、异常说明外，还必须指出该方法做什么事情，实现什么功能。

   > 说明：对子类的实现要求，或者调用注意事项，请一并说明。

3. 【强制】所有的类都必须添加创建者和创建日期。

4. 【强制】所有的枚举类型字段必须要有注释，说明每个数据项的用途。

5. 【推荐】方法内部单行注释，在被注释语句上方另起一行，使用//注释。方法内部多行注释使用/* */注释，注意与代码对齐。

6. 【推荐】与其“半吊子”英文来注释，不如用中文注释把问题说清楚。专有名词与关键字保持英文原文即可。

   > 反例：“TCP 连接超时”解释成“传输控制协议连接超时”，理解反而费脑筋。

7. 【推荐】代码修改的同时，注释也要进行相应的修改，尤其是参数、返回值、异常、核心逻辑等的修改。

   > 说明：代码与注释更新不同步，就像路网与导航软件更新不同步一样，如果导航软件严重滞后，就失去了导航的意义。

8. 【参考】谨慎注释掉代码。在上方详细说明，而不是简单地注释掉。如果无用，则删除。

   > 说明：代码被注释掉有两种可能性：
   > 1）后续会恢复此段代码逻辑。
   > 2）永久不用。前者如果没有备注信息，难以知晓注释动机。后者建议直接删掉（代码仓库保存了历史代码）。

9. 【参考】对于注释的要求：第一、能够准确反应设计思想和代码逻辑；第二、能够描述业务含义，使别的程序员能够迅速了解到代码背后的信息。完全没有注释的大段代码对于阅读者形同天书，注释是给自己看的，即使隔很长时间，也能清晰理解当时的思路；注释也是给继任者看的，使其能够快速接替自己的工作。

10. 【参考】好的命名、代码结构是自解释的，注释力求精简准确、表达到位。避免出现注释的一个极端：过多过滥的注释，代码的逻辑一旦修改，修改注释是相当大的负担。

    > 反例： put elephant into fridge put(elephant, fridge); 方法名 put，加上两个有意义的变量名 elephant 和 fridge，已经说明了这是在干什么，语义清晰的代码不需要额外的注释。

11. 【参考】特殊注释标记，请注明标记人与标记时间。注意及时处理这些标记，通过标记扫描，经常清理此类标记。线上故障有时候就是来源于这些标记处的代码。

    > 1） 待办事宜（TODO）:（ 标记人，标记时间，[预计处理时间]） 表示需要实现，但目前还未实现的功能。这实际上是一个 Javadoc 的标签，目前的 Javadoc 还没有实现，但已经被广泛使用。只能应用于类，接口和方法（因为它是一个 Javadoc 标签）。
    > 2） 错误，不能工作（FIXME）:（标记人，标记时间，[预计处理时间]） 在注释中用 FIXME 标记某代码是错误的，而且不能工作，需要及时纠正的情况。

### 内存管理

1. 【强制】静态集合类引起内存泄漏，像HashMap、Vector等的使用最容易出现内存泄露，这些静态变量的生命周期和应用程序一致，他们所引用的所有对象Object也不能被释放，因为他们也将一直被Vector等引用着，必须从Vector 中删除。

   > 说明：在这个例子中，循环申请Object 对象，并将所申请的对象放入一个Vector 中，如果仅仅释放引用本身（o=null），那么Vector 仍然引用该对象，所以这个对象对GC 来说是不可回收的。因此，如果对象加入到Vector 后，还必须从Vector 中删除，最简单的方法就是将Vector对象设置为null
   > 反例：
   >
   > ```
   >    Static Vector v = new Vector(10);    
   >    for (int i = 1; i<100; i++)
   >    {
   >        Object o = new Object();
   >        v.add(o);
   >        o = null;
   >    }
   > ```

2. 【推荐】建议手动将生成的无用对象，中间对象置为null，加快内存回收。

3. 【推荐】JVM调优通过配置JVM的参数来提高垃圾回收的速度，如果在没有出现内存泄露且上面两种办法都不能保证JVM内存回收时，可以考虑采用JVM调优的方式来解决，不过一定要经过实体机的长期测试，因为不同的参数可能引起不同的效果。给JVM环境参数设置-XX:+HeapDumpOnOutOfMemoryError参数，让JVM碰到OOM场景时输出dump信息。在线上生产环境，JVM的Xms和Xmx设置一样大小的内存容量，避免在GC 后调整堆大小带来的压力。

### 其它

1. 【强制】在使用正则表达式时，利用好其预编译功能，可以有效加快正则匹配速度。

   > 说明：不要在方法体内定义：Pattern pattern = Pattern.compile(“规则”);

2. 【强制】注意 Math.random() 这个方法返回是 double 类型，注意取值的范围 0≤x<1（能够取到零值，注意除零异常），如果想获取整数类型的随机数，不要将 x 放大 10 的若干倍然后取整，直接使用 Random 对象的 nextInt 或者 nextLong 方法。

3. 【强制】获取当前毫秒数 System.currentTimeMillis(); 而不是 new Date().getTime();

   > 说明：如果想获取更加精确的纳秒级时间值，使用 System.nanoTime()的方式。在 JDK8 中，针对统计时间等场景，推荐使用 Instant 类。

4. 【强制】RPC服务选择JRES服务支持的参数。

5. 【推荐】及时清理不再使用的代码段或配置信息。

   > 说明：对于垃圾代码或过时配置，坚决清理干净，避免程序过度臃肿，代码冗余。 正例：对于暂时被注释掉，后续可能恢复使用的代码片断，在注释代码上方，统一规定使用三个斜杠(///)来说明注释掉代码的理由。

## 二、异常日志

### 异常处理

1. 【强制】Java 类库中定义的可以通过预检查方式规避的 RuntimeException 异常不应该通过catch 的方式来处理，比如：NullPointerException，IndexOutOfBoundsException 等等。

   > 说明：无法通过预检查的异常除外，比如，在解析字符串形式的数字时，不得不通过 catch NumberFormatException 来实现。
   > 正例：if (obj != null) {...}
   > 反例：try { obj.method(); } catch (NullPointerException e) {…}

2. 【强制】异常不要用来做流程控制，条件控制。

   > 说明：异常设计的初衷是解决程序运行中的各种意外情况，且异常的处理效率比条件判断方式要低很多。

3. 【强制】捕获异常是为了处理它，不要捕获了却什么都不处理而抛弃之，如果不想处理它，请将该异常抛给它的调用者。最外层的业务使用者，必须处理异常，将其转化为用户可以理解的内容。

4. 【强制】不要在 finally 块中使用 return。

   > 说明：finally 块中的 return 返回后方法结束执行，不会再执行 try 块中的 return 语句。

5. 【推荐】方法的返回值可以为 null，不强制返回空集合，或者空对象等，必须添加注释充分说明什么情况下会返回 null 值。

   > 说明：本手册明确防止 NPE 是调用者的责任。即使被调用方法返回空集合或者空对象，对调用 者来说，也并非高枕无忧，必须考虑到远程调用失败、序列化失败、运行时异常等场景返回 null 的情况。

6. 【推荐】防止 NPE，是程序员的基本修养，注意 NPE 产生的场景：

   > 1）返回类型为基本数据类型，return 包装数据类型的对象时，自动拆箱有可能产生 NPE。
   > 反例：public int f() { return Integer 对象}， 如果为 null，自动解箱抛 NPE。
   > 2） 数据库的查询结果可能为 null。
   > 3） 集合里的元素即使 isNotEmpty，取出的数据元素也可能为 null。
   > 4） 远程调用返回对象时，一律要求进行空指针判断，防止 NPE。
   > 5） 对于 Session 中获取的数据，建议 NPE 检查，避免空指针。
   > 6） 级联调用 obj.getA().getB().getC()；一连串调用，易产生 NPE。
   > 正例：使用 JDK8 的 Optional 类来防止 NPE 问题。

7. 【参考】对于公司外的 http/api 开放接口必须使用“错误码”；而应用内部推荐异常抛出；跨应用间 RPC 调用优先考虑使用 Result 方式，封装 isSuccess()方法、“错误码”、“错误简短信息”。

   > 说明：关于 RPC 方法返回方式使用 Result 方式的理由：
   > 1）使用抛异常返回方式，调用方如果没有捕获到就会产生运行时错误。
   > 2）如果不加栈信息，只是 new 自定义异常，加入自己的理解的 error message，对于调用 端解决问题的帮助不会太多。如果加了栈信息，在频繁调用出错的情况下，数据序列化和传输 的性能损耗也是问题。

### 日志规约

1. 【强制】应用中不可直接使用日志系统（Log4j、Logback）中的 API，而应依赖使用日志框架 SLF4J 中的 API，使用门面模式的日志框架，有利于维护和各个类的日志处理方式统一。

   > import org.slf4j.Logger;
   >
   > import org.slf4j.LoggerFactory;
   >
   > private static final Logger logger = LoggerFactory.getLogger(Abc.class);

2. 【推荐】对 trace/debug/info 级别的日志输出，必须使用条件输出形式或者使用占位符的方式。

   > 说明：logger.debug("Processing trade with id: " + id + " and symbol: " + symbol); 如果日志级别是 warn，上述日志不会打印，但是会执行字符串拼接操作，如果 symbol 是对象，会执行 toString()方法，浪费了系统资源，执行了上述操作，最终日志却没有打印。
   > 正例：（条件）建设采用如下方式 if (logger.isDebugEnabled()) { logger.debug("Processing trade with id: " + id + " and symbol: " + symbol); }
   > 正例：（占位符） logger.debug("Processing trade with id: {} and symbol : {} ", id, symbol);

3. 【推荐】异常信息应该包括两类信息：案发现场信息和异常堆栈信息。如果不处理，那么通过关键字 throws 往上抛出。

   > 正例：logger.error(各类参数或者对象 toString() + "_" + e.getMessage(), e);

4. 【推荐】谨慎地记录日志。生产环境禁止输出 debug 日志；有选择地输出 info 日志；如果使用 warn 来记录刚上线时的业务行为信息，一定要注意日志输出量的问题，避免把服务器磁盘撑爆，并记得及时删除这些观察日志。

   > 说明：大量地输出无效日志，不利于系统性能提升，也不利于快速定位错误点。记录日志时请思考：这些日志真的有人看吗？看到这条日志你能做什么？能不能给问题排查带来好处？

## 三、单元测试

1. 【强制】单元测试代码必须写在如下工程目录：src/test/java，不允许写在业务代码目录下

2. 【强制】好的单元测试必须遵守 AIR 原则。

   > 说明：单元测试在线上运行时，感觉像空气（AIR）一样并不存在，但在测试质量的保障上，却是非常关键的。好的单元测试宏观上来说，具有自动化、独立性、可重复执行的特点。
   >
   > - A：Automatic（自动化）
   > - I：Independent（独立性）
   > - R：Repeatable（可重复）

3. 【推荐】单元测试的基本目标：语句覆盖率达到70%；核心模块的语句覆盖率和分支覆盖率都要达到100%

## 四、工程结构

### 应用分层

> 【推荐】参考[am-jelly-service](https://git01.hundsun.com/wusong26166/am-jelly-service.git)模板项目分层结构 

**jelly-service**

> Service 层（业务逻辑）,服务入口
> 目录说明
>
> - src/main/java 源码目录
> - src/main/resources 资源文件/配置文件目录
> - deploy 部署配置目录
> - src/test/java 单元测试代码

**jelly-service-api**

> Service 接口层，可供rpc消费端调用

**jelly-common**

> Common 层
>
> 1. 工具类
> 2. 公共类

**jelly-dao**

> DAO 层(数据持久层)

```
## DAO 层主要存在以下3个文件

# Mybatis Mapper 接口定义
src/main/java/com.hundsun.am.jelly.dao.mapper.BaseDictionaryMapper

# 自动生成的 Mybatis Mapper Xml，禁止修改。
src/generated/resources/META-INF/mybatis/mapper/BaseDictionaryMapper.generated.xml

# 手动维护的 Mybatis Mapper Xml，与 BaseDictionaryMapper.xml 对应。
src/main/resources/META-INF/mybatis/mapper/BaseDictionaryMapper.xml
```

**jelly-domain**

> 领域层 Domain 层主要存放实体类
>
> 1. 与数据库表一一对应的实体类，该类一般由`Mybatis Generator`自动生成，应该避免修改。
> 2. 联表查询时对多个数据库表的组合类
> 3. 实体类相关的枚举类
> 4. 其它 POJO 对象

**jelly-manager**

> Manager 层（业务逻辑）
>
> 1. 事务处理
> 2. 第三方接口调用
> 3. 数据聚合
> 4. etc.

### 二方库

1. 【强制】二方库里可以定义枚举类型，参数可以使用枚举类型，但是接口返回值不允许使用枚举类型或者包含枚举类型的 POJO 对象。

   > 说明：返回值禁止使用枚举类型是为了支持接口扩展后的兼容性。试想如果返回值使用了枚举类型，因业务需要对枚举类型进行了扩展，在客户端API没有升级情况下，扩展后的枚举值是无法正常序列化的。相反，枚举类型作为入参不会有兼容性问题，且还能限制值域范围，起到了安全作用。

2. 【推荐】所有 pom 文件中的依赖声明放在\<dependencies>语句块中，所有版本仲裁放在 \<dependencyManagement>语句块中。

3. 为避免应用二方库的依赖冲突问题，二方库发布者应当遵循以下原则：

   > 1）【强制】精简可控原则。移除一切不必要的 API 和依赖，只包含 ServiceAPI、必要的领域模型对象、Utils类、常量、枚举等。如果依赖其它二方库，尽量是 provided引入，让二方库使用者去依赖具体版本号。
   > 正例：jelly-service-api服务接口包满足精简化依赖，和api接口定义无关的（如mq、redis），在实现层引用；
   > 2）【推荐】稳定可追溯原则。每个版本的变化应该被记录，二方库由谁维护，源码在哪里，都需要能方便查到。除非用户主动升级版本，否则公共二方库的行为不应该发生变化。

4. 【强制】引用二方库、三方库时，禁止依赖 SNAPSHOT 版本，禁止采用 LATEST、 RELEASE 版本号依赖方式。工程依赖的编译插件(build->plugin)必须指明具体版本号。

   > 说明：不依赖 SNAPSHOT、LATEST、RELEASE 版本是保证应用发布的幂等性。 另外， 也可以加快编译时的打包构建，MAVEN 3已经移除LATEST、 RELEASE版本号依赖方式。
   > MAVEN编译插件plugin如果没有指明具体版本号且父pom没有使用dependencyManagement管理版本号，则项目默认获取的是当前插件最新版本，不能保证应用打包幂等性。

## 五、设计规约

1. 【强制】避免重复造轮子，优先选择研发中心已发布组件，特殊情况下需架构组评审；

2. 【推荐】谨慎使用继承的方式来进行扩展，优先使用聚合/组合的方式来实现。

3. 【推荐】系统架构设计的目的：

   > 1）确定系统边界。确定系统在技术层面上的做与不做。
   > 2）确定系统内模块之间的关系。确定模块之间的依赖关系及模块的宏观输入与输出。
   > 3）确定指导后续设计与演化的原则。使后续的子系统或模块设计在规定的框架内继续演化。
   > 4）确定非功能性需求。非功能性需求是指安全性、可用性、可扩展性等。

4. 【推荐】H5页面直接通过IAR访问T2（JAVA、C++）服务，不需要项目额外设计包装一层HTTP API接口

   > 说明：IAR有http转T2功能，当T2服务需要暴露给H5调用时，可通过服务治理自动生成的http接口上线或在IAR配置http地址与T2功能号对应关系。

## 六、缓存规约

框架封装了多种缓存实现，根据业务需要选择调用，一般情况选择分布式缓存，访问频率超高、有极速需求、仅升级时才会更新的缓存数据选择本地缓存。

### 分布式缓存

配置及使用说明参考研发中心缓存组件文档

Redis 是一个高性能的key-value数据库，通过socket访问到缓存服务，处理集群和分布式缓存方便，有成熟的方案。存在缓存共享、分布式部署、缓存内容很大的，建议用redis。

1. 配置信息：

   ```
   hs.cache.default.type=redis  
   hs.cache.default.mode=standalone  
   hs.cache.default.hostName= xx.xx.xx.xx  
   hs.cache.default.port=6379  
   hs.cache.default.password=
   ```

2. 调用方式

   1. 包装调用，获取缓存数据时推荐使用

      ```
      @Autowired
      private RedisCacheContainer redisCacheContainer;
      ```

      方法内调用

      ```
      List<GsTFee> list = new DefaultCacheReader<>(redisCacheContainer,  
          x -> String.format((CacheSettingManager.baseDictionary.getKey())),  
          x -> gsTFeeMapper.findAll(),  
          CacheSettingManager.baseDictionary.getExpiredSeconds()  
          ).read(new TypeReference<List<GsTFee>>() {  
          });
      ```

   2. 直接调用，更新缓存数据时推荐使用

      ```
      @Autowired
      private RedisCacheContainer redisCacheContainer;
      ```

      方法内调用

      ```
      redisCacheContainer.add(xxx);
      redisCacheContainer.remove(xxx);
      ```

### 本地缓存

Ehcache-Java的进程内缓存框架，直接在jvm虚拟机中缓存，速度快，效率高，适用于对延迟要求很高的应用。缺点也很明显，在分布式系统中不能很好解决缓存同步。通常为远程分布式缓存补充方案，作为二级缓存。 框架中实现了两种分布式更新策略

1. udp广播方式 优点：能自动发现新增、下线节点，任意节点数据变化发广播通知其他节点缓存更新。缺点：跨网段或跨机房部署时广播可能不稳定；

2. 点对点配置更新方式 优点：不受网络影响，实时通知更新。缺点：每个节点均需要配置其他节点监听缓存更新地址，上下线节点需要修改整个服务组配置；

3. 配置信息

   ```
   spring.cache.type=ehcache  
   spring.cache.ehcache.config=classpath:ehcache.xml
   ```

   ehcache.xml：

   ```
   <?xml version="1.0" encoding="UTF-8"?>
   <ehcache xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
            xsi:noNamespaceSchemaLocation="http://www.ehcache.org/ehcache.xsd">
       <diskStore path="java.io.tmpdir"/>
       <!--自动发现方式-->
       <!--<cacheManagerPeerProviderFactory-->
           <!--class="net.sf.ehcache.distribution.RMICacheManagerPeerProviderFactory"-->
           <!--properties="peerDiscovery=automatic, multicastGroupAddress=224.1.1.1,  multicastGroupPort=1000, timeToLive=64"/>-->
       <!--手动配置方式-->
       <cacheManagerPeerProviderFactory
           class="net.sf.ehcache.distribution.RMICacheManagerPeerProviderFactory"
           properties="peerDiscovery=manual, rmiUrls=//10.20.29.242:1000/baseCache"/>
       <cacheManagerPeerListenerFactory
           class="net.sf.ehcache.distribution.RMICacheManagerPeerListenerFactory"
           properties="hostName=192.168.162.96,port=1000,socketTimeoutMillis=20000"/>
   
       <defaultCache
           maxElementsInMemory="10"
           overflowToDisk="false"
           diskPersistent="false"
           timeToIdleSeconds="0"
           timeToLiveSeconds="60"
           memoryStoreEvictionPolicy="LRU">
       </defaultCache>
       <cache name="baseCache"
              maxElementsInMemory="30"
              overflowToDisk="false"
              diskPersistent="false"
              timeToIdleSeconds="0"
              timeToLiveSeconds="300"
              memoryStoreEvictionPolicy="LRU">
           <cacheEventListenerFactory class="net.sf.ehcache.distribution.RMICacheReplicatorFactory"/>
           <bootstrapCacheLoaderFactory class="net.sf.ehcache.distribution.RMIBootstrapCacheLoaderFactory"/>
       </cache>
   </ehcache>
   ```

   Cache属性说明：

   必须属性：

   　　name:缓存名称。

   　　maxElementsInMemory：缓存最大个数。

   　　eternal:对象是否永久有效，一但设置了，timeout将不起作用。

   　　overflowToDisk：当内存中对象数量达

   　　maxElementsInMemory时，Ehcache将会对象写到磁盘中。

   　　 diskSpoolBufferSizeMB：这个参数设置DiskStore（磁盘缓存）的缓存区大小。默认是30MB。每个Cache都应该有自己的一个缓冲区。

   　　maxElementsOnDisk：硬盘最大缓存个数。

   可选的属性：

   　　timeToIdleSeconds：设置对象在失效前的允许闲置时间（单位：秒）。仅当eternal=false对象不是永久有效时使用，可选属性，默认值是0，也就是可闲置时间无穷大。

   　　timeToLiveSeconds：设置对象在失效前允许存活时间（单位：秒）。最大时间介于创建时间和失效时间之间。仅当eternal=false对象不是永久有效时使用，默认是0.，也就是对象存活时间无穷大。

   　　diskPersistent：是否disk store在虚拟机启动时持久化. The default value is false.

   　　memoryStoreEvictionPolicy：当达到maxElementsInMemory限制时，Ehcache将会根据指定的策略去清理内存。默认策略是LRU（最近最少使用）。你可以设置为FIFO（先进先出）或是LFU（较少使用）。

   　　diskExpiryThreadIntervalSeconds：磁盘失效线程运行时间间隔，默认是120秒。

   　　clearOnFlush：内存数量最大时是否清除。

   缓存子元素：

   　　cacheEventListenerFactory：注册相应的的缓存监听类，用于处理缓存事件，如put,remove,update,和expire

   bootstrapCacheLoaderFactory:指定相应的BootstrapCacheLoader，用于在初始化缓存，以及自动设置。

4. 调用方式

   @Cacheable 在方法执行前先检查是否有缓存数据，如果有直接返回。如果没有数据，调用方法并将方法返回值存放在缓存当中。

   调用实例：

   ```
      @Cacheable(value = CACHE_NAME, key ="'getByIdWithEhCache:'+#id.toString()")  
      public ApiResult<UserInfo> getByIdWithEhCache(Integer id) {  
          UserInfo data = userInfoMapper.findById(id.toString());  
          return ok(data);  
      }
   ```

   @CachePut 无论怎样，都将方法的返回值更新到缓存当中。

   @CacheEvict 将一条或者多条数据从缓存中删除。

## 七、消息规约

引用自研发中心消息使用说明文档

1. 依赖项

   ```
   <dependecy>
       <groupId>com.hundsun.jrescloud</groupId>
       <artifactId>jrescloud-starter-mq-amqp</artifactId >
   </denpendecy>
   ```

2. 配置信息

**Binder**

| 配置                       | 功能                  | 说明                                           |
| :------------------------- | :-------------------- | :--------------------------------------------- |
| hs.mq.binder.name.type     | Binder类型            | 值目前只支持:rabbit                            |
| hs.mq.binder.name.host     | 主机地址,broker的地址 | 如果集群，则用逗号分开                         |
| hs.mq.binder.name.port     | Broker的端口号        | 如果集群，则用逗号分开，并且和host个数保持一致 |
| hs.mq.binder.name.username | 用户名                |                                                |
| hs.mq.binder.name.password | 密码                  |                                                |

注：[name]部分自己定义。

**Binddings**

| 配置                                          | 功能                     | 说明                                                   | 默认值 |
| :-------------------------------------------- | :----------------------- | :----------------------------------------------------- | :----- |
| hs.mq.bindings.output.instanceId.binder       | 发布实例绑定的binder     | 填写binder的name                                       | 必填   |
| hs.mq.bindings.output.instanceId.destination  | 发布实例的目的地         | 自己命名                                               | 必填   |
| hs.mq.bindings.output.instanceId.exchangeType | 发布实例的类型           | topic,fanout,direct                                    | 必填   |
| hs.mq.bindings.output.instanceId.routingKey   | 发布实例的routingkey     | 对应rabbit的含义                                       | 选填   |
| hs.mq.bindings.output.instanceId.durable      | 发布是否持久化           | false,true.如果为true，则实例化rabbit的exchange和消息  | 选填   |
| hs.mq.bindings.input.instanceId. binder       | 订阅实例的binder         | 填写binder的name                                       | 必填   |
| hs.mq.bindings.input.instanceId.destination   | 订阅实例的目的地         | 自己命名                                               | 必填   |
| hs.mq.bindings.input.instanceId.exchangeType  | 订阅实例的类型           | topic,fanout,direct                                    | 必填   |
| hs.mq.bindings.input.instanceId.routingKey    | 订阅实例的routingkey     | 对应rabbit的含义                                       | 选填   |
| hs.mq.bindings.input.instanceId.durable       | 订阅队列是否持久化       | false,true                                             | 选填   |
| hs.mq.bindings.input.instanceId.group         | 订阅组                   | 如果多个消费者配置同一个组，则多个消费者消费同一个队列 | 必填   |
| hs.mq.bindings.input.instanceId.autoAck       | 订阅实例是否自动确认     | false,true                                             | 选填   |
| hs.mq.bindings.input.instanceId.prefetch      | 订阅实例预处理个数       | 正整数                                                 | 选填   |
| hs.mq.bindings.input.instanceId.autodlq       | 订阅实例是否创建死信队列 | false,true                                             | 选填   |

1. 消息发送

   1. 注入API

      ```
      @Autowired
      private OutputExchange outputExchange；
      ```

   2. 消息发布

      ```
      outputExchange.publish(String instanceId, Message message);
      ```

2. 消息订阅

   ```
   @EnableBinding
   Public class TestListener{   
       @MessageBinding(“instanceId”)
       public void receive(Message message, Channel channel){
   
       }
   }
   ```

3. MC消息推送

   1. remoting配置

      ```
      system.dev.mode=true
      remoting.trans.charset=utf-8
      remoting.annotation.scanpackage=com.hundsun.jresplus.service;
      remoting.auto.exporter=true
      remoting.auto.exporter.scanpackage=com.hundsun.jresplus.service;
      remoting.validateObject=false
      ```

   2. Application导入配置文件

      ```
      @ImportResource(locations = {"classpath*:conf/spring/jresplus-cep-beans.xml", "classpath*:conf/*-beans.xml"})
      static {
          System.setProperty(PropertyPlaceholderConfigurer.CONFIG_LOCATION, "server.properties");
      }
      ```

   3. 构建消息参数 示例：

      ```
      IDataset dataset = DatasetService.getInstace().getDataset();
      dataset.setTotalCount(1);
      dataset.addColumn("subsys_no", DatasetColumnType.DS_INT);
      dataset.addColumn("ins_id", DatasetColumnType.DS_INT);
      dataset.addColumn("ins_modify_index", DatasetColumnType.DS_INT);
      dataset.appendRow();
      dataset.updateInt(1, 3320);
      dataset.updateInt(2, 123456);
      dataset.updateInt(3, 123456);
      Integer result = McPubUtils.publish("o45.gs.inscancel", dataset);
      ```

## 八、分布式事务规约

引用自研发中心TCC使用说明文档

1. 一个完整的TCC业务由一个主业务服务和若干个从业务服务组成，主业务服务发起并完成整个业务活动，TCC模式要求从服务提供三个接口：Try、Confirm、Cancel。

   1. Try：完成所有业务检查
      预留必须业务资源
   2. Confirm：真正执行业务
      不作任何业务检查；只使用Try阶段预留的业务资源；Confirm操作满足幂等性；
   3. Cancel：取消
      释放Try阶段预留的业务资源；Cancel操作满足幂等性；

2. 框架提供对分布式事务操作进行支持，maven依赖如下：

   ```
   <dependency>
           <groupId>com.hundsun.jrescloud</groupId>
           <artifactId>jrescloud-starter-dts</artifactId>
       </dependency>
   ```

   配置：

   ```
       #分布式任务,事务持久化依赖zookeeper
       dts.zkServers=xxx.xxx.xx.xxx:2181
       dts.recoverJob=true
   ```

3. 代码范例 接口声明 ：

   ```
   @CloudService
   public interface UserInfoService {
       @Compensable
       List<UserInfo> findAll();
   }
   ```

   服务发起方：

   ```
   @Controller
   public class IndexController {
       @CloudReference
       private UserInfoService userService;
   
       @RequestMapping(value = "/index", method = RequestMethod.GET)
       @Compensable(confirmMethod = "confirmQuery", 
   cancelMethod = "cancelQuery", 
           invocationContext = RpcInvocationContext.class, 
           transactionContextEditor = RpcTransactionContextEditor.class)
       public String index(Model model, String mode) {
           userService.findAll();
           return "index";
       }
   
       public void confirmQuery(Model model, String mode) {
           System.out.println("confirmQuery");
       }
   
       public void cancelQuery(Model model, String mode) {
           System.out.println("cancelQuery");
       }
   }
   ```

   服务参与方：

   ```
   @CloudComponent
   public class UserInfoServiceImpl implements UserInfoService {
   @Resource
       private UserInfoDao userInfoDao;
   
       @Override
       @Compensable(confirmMethod = "confirmFindAll", 
   cancelMethod = "cancelFindAll", 
   invocationContext = RpcInvocationContext.class, 
   transactionContextEditor = RpcTransactionContextEditor.class)
       public List<UserInfo> findAll() {
           return userInfoDao.findAll();
       }
   
       public void confirmFindAll() {
           System.out.println("confirmFindAll");
       }
   
       public void cancelFindAll() {
           System.out.println("cancelFindAll");
       }
   }
   ```