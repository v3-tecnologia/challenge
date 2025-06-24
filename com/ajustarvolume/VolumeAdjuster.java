package challenge.com.ajustarvolume;

public class VolumeAdjuster {
    public void
    adjustVolume(Vehicle vehicle, CustomLogger logger){
        try{
            VehicleValidator.validate(vehicle);

            logger.log("Ajustando volume para" + vehicle.getPlate() + " : " + vehicle.getVolume());
            Thread.sleep(500);
            logger.log("Volume ajustado com sucesso para: " + ": " + vehicle.getPlate());
        }catch(InterruptedException e){
            logger.log("Erro a ajustar volume para " + ": " + vehicle.getPlate() 
            + ":" + e.getMessage());
            throw new
            ProcessingException("Falha ao ajustar volume", e);
        }
    }
}
