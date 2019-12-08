
public class FinalEtc{

    void divide(){

        try{
            //division by zero produces an exception
            int c=10/0;
        }catch(Exception e){
            System.out.printf("exception caught %s \n", e.getMessage());
        }finally{
            System.out.println("executing finally ...");
        }
    }

    public static void main(String []args){
        
        FinalEtc fe = new FinalEtc();
        fe.divide();
    }
}

