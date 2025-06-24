package challenge.com.ajustarvolume;
import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.TimeUnit;

public class VehicleProcessor {
private static final int THREAD_POOL_SIZE = 4;

public void processFromFile(File file, CustomLogger logger) {
List<Vehicle> vehicles = readVehiclesFromFile(file, logger);
ExecutorService executor = Executors.newFixedThreadPool(THREAD_POOL_SIZE);

for (Vehicle vehicle : vehicles) {
executor.execute(() -> {
VolumeAdjuster adjuster = new VolumeAdjuster();
adjuster.adjustVolume(vehicle, logger);
});
}

executor.shutdown();
try {
if (!executor.awaitTermination(1, TimeUnit.MINUTES)) {
executor.shutdownNow();
}
} catch (InterruptedException e) {
executor.shutdownNow();
Thread.currentThread().interrupt();
}
}

@SuppressWarnings("unused")
private List<Vehicle> readVehiclesFromFile(File file, CustomLogger logger) {
List<Vehicle> vehicles = new ArrayList<>();

try (Scanner scanner = new Scanner(file)) {

if (scanner.hasNextLine()) {
String header = scanner.nextLine();
if (!header.equals("placa,volume")) {
logger.log("Aviso: Formato de cabeçalho inesperado no arquivo");
}
}

while (scanner.hasNextLine()) {
String line = scanner.nextLine();
String[] parts = line.split(",");
if (parts.length == 2) {
try {
String plate = parts[0].trim();
int volume = Integer.parseInt(parts[1].trim());
vehicles.add(new Vehicle(plate, volume));
} catch (NumberFormatException e) {
logger.log("Erro: Volume inválido na linha - " + line);
}
} else {
logger.log("Erro: Formato inválido na linha - " + line);
}
}
} catch (FileNotFoundException e) {
throw new ProcessingException("Arquivo não encontrado: " + file.getPath(), e);
}

return vehicles;
}
}
