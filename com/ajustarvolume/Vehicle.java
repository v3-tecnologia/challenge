package challenge.com.ajustarvolume;

public class Vehicle {
    private final String plate;
    private final int volume;

    public Vehicle(String plate, int volume){
        this.plate = plate;
        this.volume = volume;
    }

    public String getPlate(){
        return plate;
    }

    public int getVolume(){
        return volume;
    }

    @Override
    public String toString(){
        return "Vehicle{plate='" + plate + "', volume=" + volume + "}";
    }
}
