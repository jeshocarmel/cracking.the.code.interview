import java.util.List;
import java.util.ArrayList;
import java.util.Random;
import java.util.Arrays;
import java.util.stream.Collectors;


public class LambdaRandom {
    
    // ordinary approach implementation
    List<Integer> getRandomSubset(List<Integer> list){
        
        List<Integer> subset= new ArrayList<Integer>();
        Random rand= new Random();
        
        for(int item: list){
            if(rand.nextBoolean()){
                subset.add(item);   
            }
        }
        return subset;
    }
    
    // lambda approach implementation
    List<Integer> getRandomSubsetUsingLambda(List<Integer> list){
        
        Random rand= new Random();
        List<Integer> subset= list.stream().filter(k-> {return rand.nextBoolean(); }).collect(Collectors.toList());
        return subset;
    }
    
    public static void main(String args[]) {
        
        List<Integer> mylist= Arrays.asList(1,2,3,5);
        
        LambdaRandom lr= new LambdaRandom();
        
        // ordinary approach
        List<Integer> subset= lr.getRandomSubset(mylist);
        System.out.println(subset);
        
        // lambda approach
        List<Integer> subsetByLambda= lr.getRandomSubsetUsingLambda(mylist);
         System.out.println(subsetByLambda);
  
    }
}