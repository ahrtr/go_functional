# go_functional

It is just an example of how to write functional program using golang. 

This program was forked from jakecoffman/go_functional.go, but I added a new function Filter.

Note that it's easier to implement this example using other languages.  

The equivalent Ruby program is as below,
```
#!/usr/bin/ruby 
vs = [1, 2, 3, 12, 20] 
puts "#{vs.select{|item| item <10 }.map{|item| item * item}.reduce{|sum, item| sum += item}}"
```

THe equivalent Java program is as below,
```
import java.util.List;                                                           
import java.util.ArrayList;                                           
import java.util.Arrays;                                                                                                                                                                       
public class functional {   
    public static void main(String[] args) {                         
        List<Integer> list = new ArrayList<Integer>(Arrays.asList(1, 2, 3, 12, 20));   
        System.out.println(list.stream().filter(v -> v<10).map(v -> v*v).reduce(0, (v1,v2)->v1+v2)); 
    }                                                               
} 
```
