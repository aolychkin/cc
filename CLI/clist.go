package CLI

type CommandLine struct {
	OpID    int //Индентификатор операции, если требуется
	Line    string
	Mission string
}

//Общие команды
var Ls = CommandLine{Line: "ls", Mission: "Showing Files"}
var Date = CommandLine{Line: "date", Mission: "Showing current Date of system"}

//Установочные комманды
var InstallPreparation = CommandLine{Line: "sudo dnf -y install epel-release && sudo dnf -y upgrade",
	Mission: "Installing Preparation"}
var InstallPodman = CommandLine{Line: "sudo dnf -y install podman",
	Mission: "Installing Podman"}
var InstallNginx = CommandLine{Line: "sudo dnf -y install nginx && sudo systemctl enable nginx && " +
	"sudo systemctl start nginx && sudo firewall-cmd --permanent --add-service=http && " +
	"sudo firewall-cmd --permanent --add-service=https && sudo firewall-cmd --reload && " +
	"sudo mkdir /etc/nginx/sites-available && sudo mkdir /etc/nginx/sites-enabled",
	Mission: "Installing Nginx"}
var InstallSnap = CommandLine{Line: "sudo yum -y install snapd && sudo systemctl enable --now snapd.socket && " +
	"sudo ln -s /var/lib/snapd/snap /snap && sudo snap install core && sudo snap refresh core",
	Mission: "Installing Snap"}
var InstallCertbot = CommandLine{Line: "sudo snap install --classic certbot && " +
	"sudo ln -s /snap/bin/certbot /usr/bin/certbot",
	Mission: "Installing Snap"}

var Close = CommandLine{Line: "echo \"ccst\"", Mission: "Close Session"}
