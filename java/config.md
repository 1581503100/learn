
#### Spring 读取yml

```java
public class YamlUtils {
    private static final Logger logger = LogManager.getLogger(YamlUtils.class);

    public static Map<String, Object> yaml2Map(String yamlSource) {
        try {
            YamlMapFactoryBean yaml = new YamlMapFactoryBean();
            yaml.setResources(new ClassPathResource(yamlSource));
            return yaml.getObject();
        } catch (Exception e) {
            logger.error("Cannot read yaml", e);
            return null;
        }
    }

    public static Properties yaml2Properties(String yamlSource) {
        try {
            YamlPropertiesFactoryBean yaml = new YamlPropertiesFactoryBean();
            yaml.setResources(new ClassPathResource(yamlSource));
            return yaml.getObject();
        } catch (Exception e) {
            logger.error("Cannot read yaml", e);
            return null;
        }
    }
}


```
#### 读取properties 文件
```java
public class PropertiesUtils{
    public static Properties properties;
    
    
    public static void init(){
        properties = new Properties();
        try{
            properties.load(new FileInputStream());

        }catch (Exception e){
            e.printStackTrace(); 
        }
    }
    
    public static String get(String key){
        if(properties == null){
            init();
        }
        return properties.getProperties(key);
    }
}
```