package challenge.com.ajustarvolume;

public class VehicleValidator {

    public static void validate(Vehicle vehicle) {
        if (vehicle == null) {
            throw new IllegalArgumentException("Veículo não pode ser nulo");
        }

        validatePlate(vehicle.getPlate());
        validateVolume(vehicle.getVolume());
    }

    private static void validatePlate(String plate) {
        if (plate == null || plate.trim().isEmpty()) {
            throw new InvalidPlateException("A placa está vazia, argumento inválido");
        }

        // Validação padrão ou Mercosul
        if (!plate.matches("[A-Za-z]{3}[0-9][A-Za-z0-9][0-9]{2}")) {
            throw new InvalidPlateException("Formato de placa inválido: " + plate);
        }
    }

    private static void validateVolume(int volume) {
        if (volume < 0 || volume > 100) {
            throw new InvalidVolumeException("Volume deve estar entre 0 e 100");
        } else if (volume == 100) {
            System.out.println("Atenção: volume no máximo (100). Isso pode prejudicar sua audição.");
        }
    }
}
