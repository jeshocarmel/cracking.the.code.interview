import java.lang.reflect.Field;
import java.lang.reflect.Method;

// credits: https://www.geeksforgeeks.org/reflection-in-java/
class Test{

     // creating a private field 
     private String s; 
  
     // creating a public constructor 
     public Test()  {  s = "Solutions for Cracking the Code Interview"; } 
   
     // Creating a public method with no arguments 
     public void method1()  { 
         System.out.println("The string is " + s); 
     } 
   
     // Creating a public method with int as argument 
     public void method2(int n)  { 
         System.out.println("The number is " + n); 
     } 
   
     // creating a private method 
     private void method3() { 
         System.out.println("Private method invoked"); 
     } 
}

class ReflectionDemo{

    public static void main(String []args) throws Exception{
     
        Test t= new Test();

        Class cls =t.getClass();
        System.out.println(cls.getName());

        Method[] methods= cls.getMethods();
        for (Method m: methods){
            System.out.println(m.getName());
        }

        Field field = cls.getDeclaredField("s");
        field.setAccessible(true);
        field.set(t, "Changing the value of a field");

        Method method1= cls.getDeclaredMethod("method1");
        method1.invoke(t);

        Method method2 = cls.getDeclaredMethod("method2", int.class);
        method2.invoke(t,19);

        Method method3 = cls.getDeclaredMethod("method3");
        method3.setAccessible(true);
        method3.invoke(t);
    }
}