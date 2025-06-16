int deltaTime = 0;
int deltaTimeSpace = 0;
void setup() {
  pinMode(3, INPUT);
  pinMode(4, INPUT);
  Serial.begin(9600);
}

void loop() {
  if (digitalRead(3) == 1 ){
    deltaTime ++;
  }
  if (digitalRead(3) == 0){
    if (deltaTime < 5 && deltaTime != 0){
      Serial.println(".");
    } else if (deltaTime > 4){
      Serial.println("-");
    }
    deltaTime = 0;
  }
  if (digitalRead(4) == 1){
    if (deltaTimeSpace == 2){
      Serial.println(" ");
      deltaTimeSpace = 0;
    } else {
      deltaTimeSpace ++;
    }
  }
  delay(100);
}
