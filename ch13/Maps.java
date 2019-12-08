import java.util.ArrayList;
import java.util.HashMap;
import java.util.Iterator;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;
import java.util.TreeMap;

public class Maps{


    Map<Integer, Integer> myTreeMap= new TreeMap<Integer,Integer>();
    Map<Integer, Integer> myHashMap= new HashMap<Integer,Integer>();
    Map<Integer, Integer> myLinkedHashMap= new LinkedHashMap<Integer,Integer>();

    public Maps(List<Integer> mylist){

        Iterator it= mylist.iterator();

        while(it.hasNext()){
            int tmp= (Integer)it.next();
            myTreeMap.put(tmp,tmp);
            myHashMap.put(tmp,tmp);
            myLinkedHashMap.put(tmp,tmp);
        }
    }

    public static void main(String [] args){
        
        List<Integer> myList= new ArrayList<Integer>();
        myList.add(1);
        myList.add(0);
        myList.add(-1);

        Maps m = new Maps(myList);

        System.out.println(m.myTreeMap.keySet());
        System.out.println(m.myLinkedHashMap.keySet());
        System.out.println(m.myHashMap.keySet());
    }
}