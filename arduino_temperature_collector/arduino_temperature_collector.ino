// Sample Arduino MAX6675 Arduino Sketch

#include "max6675.h"

int ktcSO = 8;
int ktcCS = 9;
int ktcCLK = 10;

MAX6675 ktc(ktcCLK, ktcCS, ktcSO);

void setup()
{
  Serial.begin(9600);
  // give the MAX a little time to settle
  delay(1000);
}

void loop()
{
  Serial.println(ktc.readFahrenheit(), 2);
  // basic readout test

  //Serial.print("Deg C = ");
  //Serial.print(ktc.readCelsius());
  //Serial.print("\t Deg F = ");
  // Serial.write(ktc.readFahrenheit());
  //String output = String(ktc.readFahrenheit(), 2);
  //Serial.write(output.c_str());
  //Serial.println(output.c_str());
  //Serial.println();

  delay(1000);
}
