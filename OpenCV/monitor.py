import numpy as np
import cv2
import MySQLdb
import csv
from mss import mss
from PIL import Image


def main():
    host = "localhost"
    user = "root"
    passwd = "root"
    db = "WhoIsHome"

    with open('database.csv') as csvFile:
        csvReader = csv.reader(csvFile)
        for row in csvReader:
            host = row[0]
            user = row[1]
            passwd = row[2]
            db = row[3]

    print(host)
    print(user)

    db = MySQLdb.connect(host=host,
                        user=user,
                        passwd=passwd,
                        db=db)
    cur = db.cursor()
    cur.execute("SELECT * FROM examples")

    width = 500
    height = 500
    mon = {'top': 160, 'left': 160, 'width': width, 'height': height} # needs to be adjustable
    sct = mss()
    faceCascade = cv2.CascadeClassifier("") # need to load car classifiers here
    video = cv2.VideoCapture(0)

    while True:
        check, _ = video.read()

        img = Image.frombytes('RGB', (width,height), sct.grab(mon).rgb)

        if check:
            faces = faceCascade.detectMultiScale(np.array(img), 1.10, 5)

            for x,y,w,h in faces:
                img = cv2.rectangle(np.array(img), (x,y), (x + w, y + h), (0,0,255), 3)

            cv2.imshow("Car detector", np.array(img))
            key = cv2.waitKey(1)

            if key == ord('x'):
                break
        else:
            print("Failed check")
            break

    video.release()
    cv2.destroyAllWindows()


if __name__ == '__main__':
	try:
		main()
	except KeyboardInterrupt:
		print('Interrupted')
		try:
			import sys
			sys.exit(0)
		except SystemExit:
			import os
			os._exit(0)
