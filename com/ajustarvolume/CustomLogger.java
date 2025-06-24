package challenge.com.ajustarvolume;

import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;

public class CustomLogger {
public void log(String message) {
String timestamp = LocalDateTime.now().format(DateTimeFormatter.ISO_LOCAL_DATE_TIME);
System.out.println("[" + timestamp + "] " + message);
}
}