package webhook

type Reply struct {
	Message string `bson:"messsage"`
}

var stringreply = []string{
	"kamu mau ngapain kak?",
	"jangan galak galak ka",
	"banyakin sabar",
	"coba kamu share live lokasi kamu biar aku cek",
	"Ratu kayaknya gada deh tunggu aja ya",
}
