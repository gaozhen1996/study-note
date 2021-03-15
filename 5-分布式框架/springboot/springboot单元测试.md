最重要的是添加三个注解添加注解


```
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
 
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.SpringApplicationConfiguration;
import org.springframework.test.context.junit4.SpringJUnit4ClassRunner;
import org.springframework.test.context.web.WebAppConfiguration;
 
import com.smartlab.young.Application;
import com.smartlab.young.coreData.model.DataDetailModel;
import com.smartlab.young.coreData.service.CoreDataService;
 
 
@RunWith(SpringJUnit4ClassRunner.class) 
@SpringApplicationConfiguration(classes = Application.class) 
@WebAppConfiguration 
public class TestChange {
 
    @Autowired
    private CoreDataService coreDataService;
 
    @SuppressWarnings("unchecked")
    @Test
    public void listDataDetail() {
 

    }
 
 
 
 
}
```
