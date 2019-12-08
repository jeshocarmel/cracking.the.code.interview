
public class ReturnFromFinally{

    int divide(){

        int a=4;
        int b=0;

        try{
            int c=a/b;
            return c;
        }catch(Exception e){
            System.out.printf("exception caught %s \n", e.getMessage());
        }finally{
            System.out.println("executing finally ...");
        }
        return 0;
    }

    public static void main(String []args){
        
        ReturnFromFinally rff = new ReturnFromFinally();
        int ans= rff.divide();

        System.out.printf("answer: %d \n", ans);

    }
}

