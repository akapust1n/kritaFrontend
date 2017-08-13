package agregate

import (
	serv "kritaServers/backend/goserver/server"
	md "kritaServers/backend/goserver/server/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getFileSize(high int, low int, session *mgo.Collection) float64 {
	count := getFloat64(session.Find(bson.M{"images.size": bson.M{"$gt": high, "$lte": low}}).Count())
	return count
}
func getWidth(high int, low int, session *mgo.Collection) float64 {
	count := getFloat64(session.Find(bson.M{"images.width": bson.M{"$gt": high, "$lte": low}}).Count())
	return count
}
func getHeight(high int, low int, session *mgo.Collection) float64 {
	count := getFloat64(session.Find(bson.M{"images.height": bson.M{"$gt": high, "$lte": low}}).Count())
	return count
}
func getLayer(high int, low int, session *mgo.Collection) float64 {
	count := getFloat64(session.Find(bson.M{"images.numlayers": bson.M{"$gt": high, "$lte": low}}).Count())
	return count
}
func AgregateImageProps() {
	c := serv.Session.DB("telemetry").C("images")
	var ic md.ImageCollected

	ic.WD.L500.Count = getWidth(0, 500, c)
	ic.WD.L1000.Count = getWidth(500, 1000, c)
	ic.WD.L2000.Count = getWidth(1000, 2000, c)
	ic.WD.L4000.Count = getWidth(2000, 4000, c)
	ic.WD.L8000.Count = getWidth(4000, 8000, c)
	ic.WD.M8000.Count = getFloat64(c.Find(bson.M{"images.width": bson.M{"$gt": 8000}}).Count())

	ic.HD.L500.Count = getHeight(500, 1000, c)
	ic.HD.L1000.Count = getHeight(1000, 2000, c)
	ic.HD.L2000.Count = getHeight(1000, 2000, c)
	ic.HD.L4000.Count = getHeight(2000, 4000, c)
	ic.HD.L8000.Count = getHeight(4000, 8000, c)
	ic.HD.M8000.Count = getFloat64(c.Find(bson.M{"images.height": bson.M{"$gt": 8000}}).Count())

	ic.LD.L1.Count = getLayer(0, 1, c)
	ic.LD.L2.Count = getLayer(1, 2, c)
	ic.LD.L4.Count = getLayer(2, 4, c)
	ic.LD.L8.Count = getLayer(4, 8, c)
	ic.LD.L16.Count = getLayer(8, 16, c)
	ic.LD.L32.Count = getLayer(16, 32, c)
	ic.LD.L64.Count = getLayer(32, 64, c)
	ic.LD.M64.Count = getFloat64(c.Find(bson.M{"images.numlayers": bson.M{"$gt": 8000}}).Count())

	ic.ID.Mb1.Count = getFileSize(0, 1, c)
	ic.ID.Mb5.Count = getFileSize(1, 5, c)
	ic.ID.Mb10.Count = getFileSize(5, 10, c)
	ic.ID.Mb25.Count = getFileSize(10, 25, c)
	ic.ID.Mb50.Count = getFileSize(25, 50, c)
	ic.ID.Mb100.Count = getFileSize(50, 100, c)
	ic.ID.Mb200.Count = getFileSize(100, 200, c)
	ic.ID.Mb400.Count = getFileSize(200, 400, c)
	ic.ID.Mb800.Count = getFileSize(400, 800, c)
	ic.ID.More800.Count = getFloat64(c.Find(bson.M{"images.size": bson.M{"$gt": 800}}).Count())

	ic.CPD.RGBA.Count = getFloat64(c.Find(bson.M{"images.colorprofile": "RGB/Alpha"}).Count())
	ic.CPD.CMYK.Count = getFloat64(c.Find(bson.M{"images.colorprofile": "CMYK/Alpha"}).Count())
	ic.CPD.Grayscale.Count = getFloat64(c.Find(bson.M{"images.colorprofile": "Grayscale/Alpha"}).Count())
	ic.CPD.Lab.Count = getFloat64(c.Find(bson.M{"images.colorprofile": "L*a*b*/Alpha"}).Count())
	ic.CPD.XYZ.Count = getFloat64(c.Find(bson.M{"images.colorprofile": "XYZ/Alpha"}).Count())
	ic.CPD.YCbCr.Count = getFloat64(c.Find(bson.M{"images.colorprofile": "YCbCr/Alpha"}).Count())

	agregatedImageInfo = ic
}
