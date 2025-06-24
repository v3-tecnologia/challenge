package challenge.com.ajustarvolume;
import java.io.File;


public class Main {
    public static void main(String[] args) {
        try {
            String plate = null;
            Integer volume = null;
            String filePath = null;

           
            for (int i = 0; i < args.length; i++) {
                if ("--placa".equals(args[i]) && i + 1 < args.length) {
                    plate = args[i + 1];
                    i++;
                } else if ("--volume".equals(args[i]) && i + 1 < args.length) {
                    try {
                        volume = Integer.parseInt(args[i + 1]);
                    } catch (NumberFormatException e) {
                        System.err.println("Volume deve ser um número inteiro");
                        System.exit(1);
                    }
                    i++;
                } else if ("--arquivo".equals(args[i]) && i + 1 < args.length) {
                    filePath = args[i + 1];
                    i++;
                }
            }

            CustomLogger logger = new CustomLogger();

            if (filePath != null) {
                // Modo processamento em lote
                VehicleProcessor processor = new VehicleProcessor();
                processor.processFromFile(new File(filePath), logger);
            } else if (plate != null && volume != null) {
                // Modo veículo único
                VolumeAdjuster adjuster = new VolumeAdjuster();
                adjuster.adjustVolume(new Vehicle(plate, volume), logger);
            } else {
                System.out.println("Uso:");
                System.out.println(" Modo único veículo: java -jar ajustar_volume.jar --placa ABC1234 --volume 50");
                System.out.println(" Modo arquivo: java -jar ajustar_volume.jar --arquivo veiculos.csv");
                System.exit(1);
            }
        } catch (Exception e) {
            System.err.println("Erro: " + e.getMessage());
            System.exit(1);
        }
    }
}
