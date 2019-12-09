import java.util.Arrays;
import java.util.List;
import java.util.stream.Stream;

class Country{
    String continent;
    double population;

    public Country(String continent, double population){
        this.continent= continent;
        this.population=population;
    }

    public String getContinent(){
        return this.continent;
    }

    public double getPopulation(){
        return this.population;
    }
}

class LambdaExpressions{

    static final String EUROPE="Europe";
    static final String OCEANIA="Oceania";
    static final String ASIA="Asia";
    static final String AFRICA="Africa";
    static final String AMERICAS= "America";

    //Populate Data
    public List<Country> getCountryStats(){

        Country usa = new Country(AMERICAS, 372.2);
        Country mexico = new Country(AMERICAS, 129.2);
        Country brazil = new Country(AMERICAS, 209.3);

        Country france = new Country(EUROPE, 66.99);
        Country uk = new Country(EUROPE, 66.44);

        Country india= new Country(ASIA, 1339.0);
        Country china= new Country(ASIA, 1386.0);
        Country iran = new Country(ASIA, 81.16);
        
        Country australia = new Country(OCEANIA, 24.16);
        Country nz= new Country(OCEANIA, 4.79);

        Country kenya=new Country(AFRICA, 49.7);
        Country saf=new Country(AFRICA, 56.73);
        Country nigeria=new Country(AFRICA, 190.9);

        List<Country> allCountries= Arrays.asList(usa,mexico, brazil,france,uk,india,china,iran,australia,nz,kenya,saf,nigeria);
        return allCountries;
    }

    //Approach 1
    public double getPopulationTraditional(List<Country> countries, String continent){

        double totalpopulation=0;

        for( Country cntry: countries){
            if(cntry.continent == continent){
                totalpopulation+= cntry.population;
            }
        }
        return totalpopulation;
    }


    //Approach 2
    public double getPopulationLambda(List<Country> countries, String continent){
     
        Stream<Country> filteredCountries= countries.stream().filter(country -> {
            return country.continent.equals(continent);
        });

        Stream<Double> populations= filteredCountries.map(c-> c.getPopulation());

        double population = populations.reduce((double) 0,(a, b) -> a+b) ;
        return population;
    }

    //Approach 3
    public double getPopulationLambdaOptimized(List<Country> countries, String continent){
     
        Stream<Double> populations= countries.stream().map(c -> c.getContinent().equals(continent)? c.getPopulation():0);

        double population = populations.reduce((double) 0,(a, b) -> a+b) ;
        return population;
    }



    public static void main(String []args){


        LambdaExpressions le= new LambdaExpressions();
        List<Country> countries = le.getCountryStats();

        double ans=le.getPopulationTraditional(countries, ASIA);
        System.out.println(ans);

        double ans2= le.getPopulationLambda(countries, ASIA);
        System.out.println(ans2);

        double ans3=le.getPopulationLambdaOptimized(countries, ASIA);
        System.out.println(ans3);

    }
}