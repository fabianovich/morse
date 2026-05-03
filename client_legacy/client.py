import serial
import time
from MorseCodePy import decode

ser = serial.Serial('COM3', 9600, timeout=1)
time.sleep(2)

morse = ""
message = ""

try:
    while True:
        if ser.in_waiting > 0:
            data = ser.readline().decode('utf-8')
            if "." in data:
                morse += "."
            elif "-" in data:
                morse += "-"
            else:
                decoded_string: str = decode(morse, language='english')
                morse = ""
                message += decoded_string
            print (message + " " + morse)
            

except KeyboardInterrupt:
    pass
finally:
    ser.close()
