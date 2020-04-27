# WhoIsHome
A system designed as a learning project to monitor who is home based off a live feed from a security camera.

Purpose: not much. Just for fun and learning.

**TODO:**
- [x] OpenCV script
- [ ] Cascade classifiers
- [x] MySQL database
- [ ] Web server
- [ ] Restful API
- [ ] Installer script
- [ ] Mobile app
- [ ] Docker container

### About the project

**SUBJECT TO CHANGE!**

WhoIsHome (WIH) is a multi-program project. It has serveral parts to it. The first part is a Python script
running on a Raspberry Pi. Using OpenCV, it watches a camera feed on the screen and detects different cars in 
my driveway. Each car also has it's own cascade classifier.

This Python script relays information on the precense of each car, along with the car owner name, to a MySQL database.
Of which will be used as a source of data for a web server built with Go. The server will have a single page displaying who
is home, based on the presence of their car.

Furthermore, the server will have it's own Restful API. For use with a mobile app. Currently only an android app (Java)
is planned. But an IOS version is foreseeable. This has many possibilities. From notifications, to keeping a log.

Finally, I will be creating a bash script for a quick means of setting up the server on other systems for those who want
to have their own WIH setup.

**WIH Flow:**
Camera -> OpenCV server (Python) -> Web server (Go) -> Mobile app (Java)
