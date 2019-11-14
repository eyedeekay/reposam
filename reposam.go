package reposam

import (
	"crypto/rand"
	"golang.org/x/crypto/bcrypt"
	"log"
	//"strings"

	"github.com/eyedeekay/sam-forwarder/interface"
	"github.com/eyedeekay/sam-forwarder/tcp"
	//"github.com/jcelliott/lumber"
	//"github.com/schollz/cowyo/server"
)

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}
	b, err = bcrypt.GenerateFromPassword(b, 14)
	if err != nil {
		return nil, err
	}
	return b, nil
}

//RepoSam is a structure which automatically configured the forwarding of
//a local service to i2p over the SAM API.
type RepoSam struct {
	*samforwarder.SAMForwarder
	password string
	ServeDir string
	up       bool
}

var err error

func (f *RepoSam) GetType() string {
	return "reposam"
}

func (f *RepoSam) ServeParent() {
	log.Println("Starting eepsite server", f.Base32())
	if err = f.SAMForwarder.Serve(); err != nil {
		f.Cleanup()
	}
}

//Serve starts the SAM connection and and forwards the local host:port to i2p
func (f *RepoSam) Serve() error {
	go f.ServeParent()
	if f.Up() {
		/*sec, err := GenerateRandomBytes(256)
		if err != nil {
			return err
		}
		hostport := strings.SplitN(f.Target(), ":", 2)
		log.Println("Starting web server", f.Target())*/
		/*server.Serve(
			f.ServeDir,
			hostport[0],
			hostport[1],
			//c.GlobalString("cert"),
			//c.GlobalString("key"),
			//TLS,
			"",
			"",
			false,
			//c.GlobalString("css"),
			//c.GlobalString("default-page"),
			"",
			"",
			//c.GlobalString("lock"),
			"",
			//c.GlobalInt("debounce"),
			5000,
			//c.GlobalBool("diary"),
			true,
			string(sec),
			//c.GlobalString("access-code"),
			f.password,
			//c.GlobalBool("allow-insecure-markup"),
			false,
			//c.GlobalBool("allow-file-uploads"),
			//c.GlobalUint("max-upload-mb"),
			false,
			0,
			//c.GlobalUint("max-document-length"),
			100000000,
			logger(false),
		)*/
	}
	return nil
}

/*func logger(debug bool) *lumber.ConsoleLogger {
	if !debug {
		return lumber.NewConsoleLogger(lumber.WARN)
	}
	return lumber.NewConsoleLogger(lumber.TRACE)

}*/

func (f *RepoSam) Up() bool {
	return f.up
}

//Close shuts the whole thing down.
func (f *RepoSam) Close() error {
	return f.SAMForwarder.Close()
}

func (s *RepoSam) Load() (samtunnel.SAMTunnel, error) {
	if !s.up {
		log.Println("Started putting tunnel up")
	}
	f, e := s.SAMForwarder.Load()
	if e != nil {
		return nil, e
	}
	s.SAMForwarder = f.(*samforwarder.SAMForwarder)
	s.up = true
	log.Println("Finished putting tunnel up")
	return s, nil
}

//NewRepoSam makes a new SAM forwarder with default options, accepts host:port arguments
func NewRepoSam(host, port string) (*RepoSam, error) {
	return NewRepoSamFromOptions(SetHost(host), SetPort(port))
}

//NewRepoSamFromOptions makes a new SAM forwarder with default options, accepts host:port arguments
func NewRepoSamFromOptions(opts ...func(*RepoSam) error) (*RepoSam, error) {
	var s RepoSam
	s.SAMForwarder = &samforwarder.SAMForwarder{}
	log.Println("Initializing eephttpd")
	for _, o := range opts {
		if err := o(&s); err != nil {
			return nil, err
		}
	}
	s.SAMForwarder.Config().SaveFile = true
	l, e := s.Load()
	//log.Println("Options loaded", s.Print())
	if e != nil {
		return nil, e
	}
	return l.(*RepoSam), nil
}
