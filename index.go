package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	// "golang.org/x/crypto/bcrypt"

	// "github.com/biezhi/gorm-paginator/pagination"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func initialMigration() {
	//open a db connection
	// var err error
	// db, err := gorm.Open("postgres", "root:@/gomanagement?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("postgres", "user=postgres dbname=mydb password=19121997 sslmode=disable")
	db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8 auto_increment=1")
	if err != nil {
		panic("failed to connect database")
	}
	//Migrate the schema
	db.AutoMigrate(&userModel{})
}

var db *gorm.DB

// func main() {
// 	fmt.Println("cat")
// 	router := gin.Default()
// 	router.GET("/", allUser)
// 	router.POST("/", createUser)
// 	// router.PUT("/:id", updateUser)
// 	// router.DELETE("/:id", deleteUser)
// 	router.Run(":3000")
// }

type userModel struct {
	gorm.Model
	// Id      string `json:"id"`
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
	Dob      string `json:"dob" gorm:"not null"`
	Phone    string `json:"phone" gorm:"not null"`
	Level    string `json:"level" gorm:"default:'User'"`
}
type loginn struct {
	email    string
	password string
}

type Dashboard struct {
	UserNumber int `json:"userNumber`
	NvtNumber  int `json:"nvtNumber`
	CveNumber  int `json:"cveNumber`
	CpeNumber  int `json:"cpeNumber`
}

type User struct {
	Id          int    `json:"id"`
	Uuid        string `json:"uuid"`
	Name        string `json:"name"`
	Comment     string `json:"comment"`
	Role        string `json:"role"`
	Password    string `json:"password"`
	Host_allow  string `json:"host_allow"`
	Hosts       string `json:"hosts"`
	Iface_allow string `json:"iface_allow"`
	Ifaces      string `json:"ifaces"`
	Created     string `json:"created"`
	Modified    string `json:"modified"`
}

type Target_Task struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type PortList struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Target struct {
	Id            int           `json:"id"`
	Uuid          string        `json:"uuid"`
	Name          string        `json:"name"`
	Hosts         string        `json:"hosts"`
	PortList      string        `json:"portlist"`
	Comment       string        `json:"comment"`
	MaxNumberHost string        `json:"maxhost"`
	RLOnly        string        `json:"rlonly"`
	RLUnify       string        `json:"rlunify"`
	Created       string        `json:"created"`
	Modified      string        `json:"modified"`
	AliveTest     string        `json:"alivetest"`
	Task          []Target_Task `json:"task"`
}

type Task struct {
	Id       int    `json:"id"`
	Uuid     string `json:"uuid"`
	Name     string `json:"name"`
	Status   string `json:"status"`
	Reports  string `json:"report"`
	Severity string `json:"severity"`
}

type NvtInfo struct {
	Cvss_vector   string `json:"cvss_vector"`
	Qod_type      string `json:"qod_type"`
	Summary       string `json:"summary"`
	Vuldetect     string `json:"vuldetect"`
	Insight       string `json:"insight"`
	Impact        string `json:"impact"`
	Affected      string `json:"affected"`
	Solution      string `json:"solution"`
	Solution_type string `json:"solution_type"`
}

type Nvt struct {
	Id        int      `json:"id"`
	Uuid      string   `json:"uuid"`
	Name      string   `json:"name"`
	Family    string   `json:"family"`
	Created   string   `json:"created"`
	Modified  string   `json:"modified"`
	Cve       []string `json:"cve"`
	Cvss_base float64  `json:"severity"`
	Qod       string   `json:"qod"`
	Xref      []string `json:"xref"`
	Tag       NvtInfo  `json:"tag"`
}

type Cve_Nvt struct {
	Nvt_Id   string `json:"nvt_id"`
	Nvt_Name string `json:"nvt_name"`
}
type VulnerableProducts struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type Cve struct {
	Id                    int                  `json:"id"`
	Name                  string               `json:"name"`
	Description           string               `json:"description"`
	Vector                string               `json:"vector"`
	Complexity            string               `json:"complexity"`
	Authentication        string               `json:"authentication"`
	ConfidentialityImpact string               `json:"confidentiality_impact"`
	IntegrityImpact       string               `json:"integrity_impact"`
	AvailabilityImpact    string               `json:"availability_impact"`
	Published             string               `json:"published"`
	Modified              string               `json:"modified"`
	Product               []VulnerableProducts `json:"vulnerableProduct"`
	Nvt                   []Cve_Nvt            `json:"nvt"`
	Severity              sql.NullString       `json:"severity"`
}

type ReportedVulnerabilites struct {
	Cve      string `json:"cve"`
	Severity string `json:"severity"`
}
type Cpe struct {
	Id         int                      `json:"id"`
	Name       string                   `json:"name"`
	Title      sql.NullString           `json:"title"`
	Created    string                   `json:"created"`
	Modified   string                   `json:"modified"`
	Cves       int                      `json:"cves"`
	Status     sql.NullString           `json:"status"`
	Severity   sql.NullString           `json:"severity"`
	Nvd_id     sql.NullInt64            `json:"nvd_id"`
	Cve_Render []ReportedVulnerabilites `json:"reportedVulnerabilites"`
}
type Sources struct {
	Report string `json:"report"`
	Nvt    string `json:"nvt"`
}

type Identifiers struct {
	Uuid    string  `json:"uuid"`
	Name    string  `json:"name"`
	Value   string  `json:"value"`
	Created string  `json:"created"`
	Source  Sources `json:"source"`
}

type Host struct {
	Id         int             `json:"id"`
	Uuid       string          `json:"uuid"`
	Name       string          `json:"name"`
	Comment    string          `json:"comment"`
	Hostname   sql.NullString  `json:"hostname"`
	IpAddress  sql.NullString  `json:"ipaddress"`
	Severity   sql.NullFloat64 `json:"severity"`
	Modified   string          `json:"modified"`
	Identifier []Identifiers   `json:"identifier"`
}

type Paginator struct {
	TotalRecord int         `json:"total_record"`
	TotalPage   int         `json:"total_page"`
	Records     interface{} `json:"records"`
	Offset      int         `json:"offset"`
	Limit       int         `json:"limit"`
	Page        int         `json:"page"`
	PrevPage    int         `json:"prev_page"`
	NextPage    int         `json:"next_page"`
}

func getDashboard(w http.ResponseWriter, r *http.Request) {
	var dashboard Dashboard
	// var dashboards []Dashboard
	db.Raw("Select count(*) from users").Scan(&dashboard.UserNumber)
	db.Raw("Select count(*) from nvts").Scan(&dashboard.NvtNumber)
	db.Raw("Select count(*) from scap.cves").Scan(&dashboard.CveNumber)
	db.Raw("Select count(*) from scap.cpes").Scan(&dashboard.CpeNumber)
	// dashboards = append(dashboards, dashboard)
	json.NewEncoder(w).Encode(dashboard)
}

func allTargets(w http.ResponseWriter, r *http.Request) {

	// var paginator Paginator
	var targets []Target
	vars := mux.Vars(r)
	page, err := strconv.Atoi(vars["page"])

	var offset int
	limit := 10
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * limit
	}

	rows, err := db.Raw("SELECT t.id, t.uuid, t.name, t.hosts, p.name FROM targets t, port_lists p WHERE t.port_list = p.id LIMIT ? OFFSET ?", limit, offset).Rows()
	if err != nil {
		log.Print(err)
		return
	}

	for rows.Next() {
		var target Target
		err = rows.Scan(&target.Id, &target.Uuid, &target.Name, &target.Hosts, &target.PortList)
		if err != nil {
			log.Print(err)
			return
		}
		targets = append(targets, target)
	}

	var count int
	row := db.Raw("Select count(*) from targets").Row() // (*sql.Row)
	row.Scan(&count)
	paginator := Paging(page, limit, offset, count, &targets)
	json.NewEncoder(w).Encode(paginator)
}

func getTarget(w http.ResponseWriter, r *http.Request) {
	// db, err := gorm.Open("postgres", "host=112.137.129.225 user=postgres dbname=gvmd password= sslmode=disable")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	var targets []Target
	vars := mux.Vars(r)
	targetId := vars["id"]
	// cpeName := vars["name"]
	rows, err := db.Raw("SELECT t.id, t.uuid, t.name, t.hosts, t.reverse_lookup_only, t.reverse_lookup_unify, t.creation_time, t.modification_time,t.alive_test, p.name as port from targets t inner join port_lists p on t.port_list = p.id where t.uuid=?", targetId).Rows()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var target Target
		err = rows.Scan(&target.Id, &target.Uuid, &target.Name, &target.Hosts, &target.RLOnly, &target.RLUnify, &target.Created, &target.Modified, &target.AliveTest, &target.PortList)
		if err != nil {
			log.Print(err)
			return
		}
		i, err := strconv.ParseInt(target.Created, 10, 64)
		j, err := strconv.ParseInt(target.Modified, 10, 64)
		if err != nil {
			panic(err)
		}
		target.Created = time.Unix(i, 0).Format(time.RFC850)
		target.Modified = time.Unix(j, 0).Format(time.RFC850)
		if target.RLOnly == "0" {
			target.RLOnly = "No"
		} else {
			target.RLOnly = "Yes"
		}

		if target.RLUnify == "0" {
			target.RLUnify = "No"
		} else {
			target.RLUnify = "Yes"
		}
		switch target.AliveTest {
		case "0":
			target.AliveTest = "Scan Config Default"
		case "1":
			target.AliveTest = "ICMP Ping"
		case "2":
			target.AliveTest = "TCP-ACK Service Ping"
		case "3":
			target.AliveTest = "TCP-SYN Service Ping"
		case "4":
			target.AliveTest = "ARP Ping"
		case "5":
			target.AliveTest = "ICMP & TCP-ACK Service Ping"
		case "6":
			target.AliveTest = "ICMP & ARP Ping"
		case "7":
			target.AliveTest = "TCP-ACK Service & ARP Ping"
		case "8":
			target.AliveTest = "ICMP, TCP-ACK Service & ARP Ping"
		case "9":
			target.AliveTest = "Consider Alive"
		}

		rows, err := db.Raw("select tasks.uuid, tasks.name from targets left join tasks on targets.id = tasks.target where targets.uuid = ? and tasks.hidden=0", targetId).Rows()
		if err != nil {
			log.Fatal(err)
		}
		var task Target_Task
		var id string
		var name string
		for rows.Next() {
			err = rows.Scan(&id, &name)
			if err != nil {
				log.Print(err)
				return
			}
			task.Id = id
			task.Name = name
			target.Task = append(target.Task, task)
		}

		targets = append(targets, target)
	}
	json.NewEncoder(w).Encode(targets)

}

func createTarget(w http.ResponseWriter, r *http.Request) {
	var target Target
	_ = json.NewDecoder(r.Body).Decode(&target)
	var u = uuid.Must(uuid.NewV4())
	var created = time.Now().Unix()
	var modified = time.Now().Unix()

	db.Exec("INSERT INTO targets(uuid,name, hosts, reverse_lookup_only, reverse_lookup_unify, comment, port_list, alive_test, creation_time, modification_time) VALUES (?,?,?,?,?,?,?,?,?,?)", &u, &target.Name, &target.Hosts, &target.RLOnly, &target.RLUnify, &target.Comment, &target.PortList, &target.AliveTest, &created, &modified)
	// if err != nil {
	// 	log.Print(err)
	// 	return
	// } else {
	// 	fmt.Fprintf(w, "Create Successful")
	// }

	// db.Create(&target)

}

func updateTarget(w http.ResponseWriter, r *http.Request) {
	// db, err := gorm.Open("postgres", "host=112.137.129.225 user=postgres dbname=gvmd password= sslmode=disable")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	vars := mux.Vars(r)
	targetId := vars["id"]

	var target Target
	_ = json.NewDecoder(r.Body).Decode(&target)
	var modified = time.Now().Unix()

	db.Exec("UPDATE targets SET name = ?, hosts = ?, reverse_lookup_only = ?, reverse_lookup_unify= ?, comment = ?, port_list = ?, alive_test = ?, modification_time = ? WHERE id = ?", &target.Name, &target.Hosts, &target.RLOnly, &target.RLUnify, &target.Comment, &target.PortList, &target.AliveTest, &modified, &targetId)

}

func deleteTarget(w http.ResponseWriter, r *http.Request) {
	// db, err := gorm.Open("postgres", "host=112.137.129.225 user=postgres dbname=gvmd password= sslmode=disable")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	vars := mux.Vars(r)
	targetId := vars["id"]

	db.Exec("INSERT INTO targets_trash(uuid, owner, name, hosts, exclude_hosts, reverse_lookup_only, reverse_lookup_unify, comment, port_list, alive_test, creation_time, modification_time) SELECT uuid, owner, name, hosts, exclude_hosts, reverse_lookup_only, reverse_lookup_unify, comment, port_list, alive_test, creation_time, modification_time FROM targets WHERE id = ?", &targetId)
	db.Exec("DELETE FROM targets WHERE id = ?", &targetId)
}

func allPortList(w http.ResponseWriter, r *http.Request) {
	// db, err := gorm.Open("postgres", "host=112.137.129.225 user=postgres dbname=gvmd password= sslmode=disable")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	var portLists []PortList
	rows, err := db.Raw("SELECT id, name FROM port_lists").Rows()
	if err != nil {
		log.Print(err)
		return
	}
	var portList PortList
	for rows.Next() {
		err = rows.Scan(&portList.Id, &portList.Name)
		if err != nil {
			log.Print(err)
			return
		}
		portLists = append(portLists, portList)
	}
	json.NewEncoder(w).Encode(portLists)
}

func allTasks(w http.ResponseWriter, r *http.Request) {
	// db, err := gorm.Open("postgres", "host=112.137.129.225 user=postgres dbname=gvmd password= sslmode=disable")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// var paginator Paginator
	var tasks []Task
	vars := mux.Vars(r)
	page, err := strconv.Atoi(vars["page"])

	var offset int
	limit := 10
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * limit
	}

	rows, err := db.Raw("SELECT id, uuid, name FROM tasks WHERE hidden = 0 LIMIT ? OFFSET ?", limit, offset).Rows()
	if err != nil {
		log.Print(err)
		return
	}

	for rows.Next() {
		var task Task
		err = rows.Scan(&task.Id, &task.Uuid, &task.Name)
		if err != nil {
			log.Print(err)
			return
		}
		tasks = append(tasks, task)
	}

	var count int
	row := db.Raw("Select count(*) from tasks where hidden = 0").Row() // (*sql.Row)
	row.Scan(&count)
	paginator := Paging(page, limit, offset, count, &tasks)
	json.NewEncoder(w).Encode(paginator)
}

func allNvts(w http.ResponseWriter, r *http.Request) {
	// db, err := gorm.Open("postgres", "host=112.137.129.225 user=postgres dbname=gvmd password= sslmode=disable")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// var paginator Paginator
	var nvts []Nvt
	vars := mux.Vars(r)
	page, err := strconv.Atoi(vars["page"])

	var offset int
	limit := 10
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * limit
	}

	rows, err := db.Raw("SELECT id, uuid, name, family, creation_time, modification_time, cve, cvss_base, qod FROM nvts LIMIT ? OFFSET ?", limit, offset).Rows()
	if err != nil {
		log.Print(err)
		return
	}

	for rows.Next() {
		var nvt Nvt
		var cve string
		err = rows.Scan(&nvt.Id, &nvt.Uuid, &nvt.Name, &nvt.Family, &nvt.Created, &nvt.Modified, &cve, &nvt.Cvss_base, &nvt.Qod)
		if err != nil {
			log.Print(err)
			return
		}
		c := " %"
		// fmt.Println(s + string(c));
		nvt.Qod = nvt.Qod + string(c)
		i, err := strconv.ParseInt(nvt.Created, 10, 64)
		j, err := strconv.ParseInt(nvt.Modified, 10, 64)
		if err != nil {
			panic(err)
		}
		nvt.Created = time.Unix(i, 0).Format(time.RFC850)
		nvt.Modified = time.Unix(j, 0).Format(time.RFC850)
		// nvt.Creation_time = time.Unix(nvt.Created, 0).Format(time.RFC850)
		// // nvt.Creation_time.Format(time.RFC850)
		// // nvt.Creation_time = nvt.Created.Format(time.RFC850))
		// nvt.Modification_time = time.Unix(nvt.Modified, 0).Format(time.RFC850)
		nvt.Cve = strings.Split(cve, ",")
		for i := range nvt.Cve {
			nvt.Cve[i] = strings.TrimSpace(nvt.Cve[i])
			if nvt.Cve[i] == "NOCVE" {
				nvt.Cve[i] = ""
			}
		}
		nvts = append(nvts, nvt)
	}

	var count int
	row := db.Raw("Select count(*) from nvts").Row() // (*sql.Row)
	row.Scan(&count)
	paginator := Paging(page, limit, offset, count, &nvts)
	json.NewEncoder(w).Encode(paginator)
}

func getNvt(w http.ResponseWriter, r *http.Request) {
	// db, err := gorm.Open("postgres", "host=112.137.129.225 user=postgres dbname=gvmd password= sslmode=disable")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	vars := mux.Vars(r)
	nvtId := vars["id"]
	// cpeName := vars["name"]
	rows, err := db.Raw("SELECT uuid, name, family, creation_time, modification_time, cve, cvss_base, qod, xref, tag from nvts where uuid =?", nvtId).Rows()
	if err != nil {
		log.Fatal(err)
	}
	var nvts []Nvt
	var cve string
	var xref string
	var tag string
	for rows.Next() {
		var nvt Nvt

		err = rows.Scan(&nvt.Uuid, &nvt.Name, &nvt.Family, &nvt.Created, &nvt.Modified, &cve, &nvt.Cvss_base, &nvt.Qod, &xref, &tag)
		if err != nil {
			log.Print(err)
			return
		}
		i, err := strconv.ParseInt(nvt.Created, 10, 64)
		j, err := strconv.ParseInt(nvt.Modified, 10, 64)
		if err != nil {
			panic(err)
		}
		nvt.Created = time.Unix(i, 0).Format(time.RFC850)
		nvt.Modified = time.Unix(j, 0).Format(time.RFC850)
		nvt.Cve = strings.Split(cve, ",")
		for i := range nvt.Cve {
			nvt.Cve[i] = strings.TrimSpace(nvt.Cve[i])
			if nvt.Cve[i] == "NOCVE" {
				nvt.Cve[i] = ""
			}
		}
		nvt.Xref = strings.Split(xref, ",")
		for i := range nvt.Xref {
			nvt.Xref[i] = strings.TrimSpace(nvt.Xref[i])
			nvt.Xref[i] = strings.TrimPrefix(nvt.Xref[i], "URL:")
		}

		var info []string
		var subinfo []string
		var nvtinfo NvtInfo
		info = strings.Split(tag, "|")
		for i := range info {
			subinfo = strings.Split(info[i], "=")

			switch subinfo[0] {
			case "cvss_base_vector":
				nvtinfo.Cvss_vector = subinfo[1]
			case "qod_type":
				nvtinfo.Qod_type = subinfo[1]
			case "summary":
				nvtinfo.Summary = subinfo[1]
			case "vuldetect":
				nvtinfo.Vuldetect = subinfo[1]
			case "insight":
				nvtinfo.Insight = subinfo[1]
			case "impact":
				nvtinfo.Impact = subinfo[1]
			case "affected":
				nvtinfo.Affected = subinfo[1]
			case "solution":
				nvtinfo.Solution = subinfo[1]
			case "solution_type":
				nvtinfo.Solution_type = subinfo[1]
			}
		}
		nvt.Tag = nvtinfo
		nvts = append(nvts, nvt)
	}
	json.NewEncoder(w).Encode(nvts)
}

func allCves(w http.ResponseWriter, r *http.Request) {
	// db, err := gorm.Open("postgres", "host=112.137.129.225 user=postgres dbname=gvmd password= sslmode=disable")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// var paginator Paginatorr
	var cves []Cve
	vars := mux.Vars(r)
	page, err := strconv.Atoi(vars["page"])

	var offset int
	limit := 10
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * limit
	}

	rows, err := db.Raw("SELECT id, name, description, vector, complexity, authentication,  confidentiality_impact, integrity_impact, availability_impact, creation_time, cvss FROM scap.cves LIMIT ? OFFSET ?", limit, offset).Rows()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var cve Cve
		err = rows.Scan(&cve.Id, &cve.Name, &cve.Description, &cve.Vector, &cve.Complexity, &cve.Authentication, &cve.ConfidentialityImpact, &cve.IntegrityImpact, &cve.AvailabilityImpact, &cve.Published, &cve.Severity)
		if err != nil {
			log.Print(err)
			return
		}
		i, err := strconv.ParseInt(cve.Published, 10, 64)
		if err != nil {
			panic(err)
		}
		cve.Published = time.Unix(i, 0).Format(time.RFC850)
		if cve.Vector == "" {
			cve.Vector = "N/A"
		}
		if cve.Complexity == "" {
			cve.Complexity = "N/A"
		}
		if cve.Authentication == "" {
			cve.Authentication = "N/A"
		}
		if cve.ConfidentialityImpact == "" {
			cve.ConfidentialityImpact = "N/A"
		}
		if cve.IntegrityImpact == "" {
			cve.IntegrityImpact = "N/A"
		}
		if cve.AvailabilityImpact == "" {
			cve.AvailabilityImpact = "N/A"
		}
		if cve.Severity.Valid == false {
			cve.Severity.String = "N/A"
		}

		cves = append(cves, cve)
	}

	var count int
	row := db.Raw("Select count(*) from scap.cves").Row() // (*sql.Row)
	row.Scan(&count)
	paginator := Paging(page, limit, offset, count, &cves)
	json.NewEncoder(w).Encode(paginator)
}

func getCve(w http.ResponseWriter, r *http.Request) {
	// db, err := gorm.Open("postgres", "host=112.137.129.225 user=postgres dbname=gvmd password= sslmode=disable")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	var product string
	vars := mux.Vars(r)
	cveName := vars["name"]
	rows, err := db.Raw("SELECT id, name, description, vector, complexity, authentication,  confidentiality_impact, integrity_impact, availability_impact, creation_time, modification_time, products, cvss FROM scap.cves where name = ?", cveName).Rows()
	if err != nil {
		log.Fatal(err)
	}
	var cves []Cve
	for rows.Next() {
		var cve Cve

		err = rows.Scan(&cve.Id, &cve.Name, &cve.Description, &cve.Vector, &cve.Complexity, &cve.Authentication, &cve.ConfidentialityImpact, &cve.IntegrityImpact, &cve.AvailabilityImpact, &cve.Published, &cve.Modified, &product, &cve.Severity)
		if err != nil {
			log.Print(err)
			return
		}
		i, err := strconv.ParseInt(cve.Published, 10, 64)
		j, err := strconv.ParseInt(cve.Modified, 10, 64)
		if err != nil {
			panic(err)
		}
		cve.Published = time.Unix(i, 0).Format(time.RFC850)
		cve.Modified = time.Unix(j, 0).Format(time.RFC850)
		if cve.Severity.Valid == false {
			cve.Severity.String = "N/A"
		}
		// cve.Product = strings.Split(product, " ")
		// for i := range cve.Product {
		// 	cve.Product[i] = strings.TrimSpace(cve.Product[i])

		// }

		rowss, err := db.Raw("select n.oid , n.name from scap.cves c inner join (select nc.oid,nc.cve_name, n.name from nvt_cves nc inner join nvts n on nc.nvt = n.id) as n on c.name = n.cve_name where c.name = ?", cveName).Rows()
		if err != nil {
			log.Fatal(err)
		}
		var cve_nvt Cve_Nvt
		var nvt_id string
		var nvt_name string
		for rowss.Next() {
			err = rowss.Scan(&nvt_id, &nvt_name)
			if err != nil {
				log.Print(err)
				return
			}
			cve_nvt.Nvt_Id = nvt_id
			cve_nvt.Nvt_Name = nvt_name
			cve.Nvt = append(cve.Nvt, cve_nvt)
		}
		//get Vulnerable Products
		rows, err := db.Raw("select a.cpe, b.name from (select c.id, c.name, ap.cpe from scap.cves c left join scap.affected_products ap on c.id = ap.cve where c.name = ?) as a inner join scap.cpes b on a.cpe = b.id", cveName).Rows()
		if err != nil {
			log.Fatal(err)
		}
		var vp VulnerableProducts
		var id int
		var name string
		for rows.Next() {
			err = rows.Scan(&id, &name)
			if err != nil {
				log.Print(err)
				return
			}
			vp.Id = id
			vp.Name = name
			cve.Product = append(cve.Product, vp)
		}

		cves = append(cves, cve)
	}
	json.NewEncoder(w).Encode(cves)

}

func allCpes(w http.ResponseWriter, r *http.Request) {
	// db, err := gorm.Open("postgres", "host=112.137.129.225 user=postgres dbname=gvmd password= sslmode=disable")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	var cpes []Cpe
	vars := mux.Vars(r)
	page, err := strconv.Atoi(vars["page"])

	var offset int
	limit := 10
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * limit
	}

	rows, err := db.Raw("SELECT id, name, title, modification_time, cve_refs, max_cvss FROM scap.cpes LIMIT ? OFFSET ?", limit, offset).Rows()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var cpe Cpe
		err = rows.Scan(&cpe.Id, &cpe.Name, &cpe.Title, &cpe.Modified, &cpe.Cves, &cpe.Severity)
		if err != nil {
			log.Print(err)
			return
		}
		i, err := strconv.ParseInt(cpe.Modified, 10, 64)
		if err != nil {
			panic(err)
		}
		cpe.Modified = time.Unix(i, 0).Format(time.RFC850)

		if cpe.Title.Valid == false {
			cpe.Title.String = "N/A"
		}
		if cpe.Severity.Valid == false {
			cpe.Severity.String = "N/A"
		}

		cpes = append(cpes, cpe)
	}

	var count int
	row := db.Raw("Select count(*) from scap.cpes").Row() // (*sql.Row)
	row.Scan(&count)
	paginator := Paging(page, limit, offset, count, &cpes)
	json.NewEncoder(w).Encode(paginator)
}

func getCpe(w http.ResponseWriter, r *http.Request) {
	// db, err := gorm.Open("postgres", "host=112.137.129.225 user=postgres dbname=gvmd password= sslmode=disable")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	vars := mux.Vars(r)
	cpeId, err := strconv.Atoi(vars["id"])
	// cpeName := vars["name"]
	rows, err := db.Raw("SELECT id, name, title, status, max_cvss, creation_time, modification_time, nvd_id FROM scap.cpes where id = ?", cpeId).Rows()
	if err != nil {
		log.Fatal(err)
	}
	var cpes []Cpe
	for rows.Next() {
		var cpe Cpe

		err = rows.Scan(&cpe.Id, &cpe.Name, &cpe.Title, &cpe.Status, &cpe.Severity, &cpe.Created, &cpe.Modified, &cpe.Nvd_id)
		if err != nil {
			log.Print(err)
			return
		}
		i, err := strconv.ParseInt(cpe.Created, 10, 64)
		j, err := strconv.ParseInt(cpe.Modified, 10, 64)
		if err != nil {
			panic(err)
		}
		cpe.Created = time.Unix(i, 0).Format(time.RFC850)
		cpe.Modified = time.Unix(j, 0).Format(time.RFC850)
		if cpe.Severity.Valid == false {
			cpe.Severity.String = "N/A"
		}
		if cpe.Status.Valid == false {
			cpe.Status.String = ""
		}
		// if cpe.Nvd_id.Valid == false {
		// 	cpe.Nvd_id.Int64
		// }

		rows, err := db.Raw("select c.name, c.cvss from scap.affected_products ap inner join scap.cves c on ap.cve = c.id where ap.cpe = ?", cpeId).Rows()
		if err != nil {
			log.Fatal(err)
		}
		var rv ReportedVulnerabilites
		var name string
		var severity string
		for rows.Next() {
			err = rows.Scan(&name, &severity)
			if err != nil {
				log.Print(err)
				return
			}
			rv.Cve = name
			rv.Severity = severity
			cpe.Cve_Render = append(cpe.Cve_Render, rv)
		}
		cpes = append(cpes, cpe)
	}
	json.NewEncoder(w).Encode(cpes)

}

func allHosts(w http.ResponseWriter, r *http.Request) {
	// db, err := gorm.Open("postgres", "host=112.137.129.225 user=postgres dbname=gvmd password= sslmode=disable")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	var hosts []Host
	vars := mux.Vars(r)
	page, err := strconv.Atoi(vars["page"])

	var offset int
	limit := 10
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * limit
	}
	rows, err := db.Raw("select h.id, h.uuid, h.name, hostname.value hostname, ip.value ip, h.modification_time, hs.severity from hosts h left join (select distinct hi.host, hi.value from (select host, max(modification_time) as modification_time from host_identifiers where name='hostname' group by host) as hos inner join host_identifiers as hi on hi.host = hos.host and hi.modification_time = hos.modification_time where hi.name='hostname' and hi.value NOT LIKE 'ww%') AS hostname on h.id = hostname.host left join (select hi.host, hi.value from (select host, max(modification_time) as modification_time from host_identifiers where name='ip' group by host) as hos inner join host_identifiers as hi on hi.host = hos.host and hi.modification_time = hos.modification_time where hi.name='ip') AS ip on h.id = ip.host left join (select distinct hs.host, hs.severity from (select host, MAX(creation_time) as creation_time from host_max_severities group by host) as hms inner join host_max_severities hs on hs.host = hms.host and hs.creation_time = hms.creation_time) as hs ON h.id = hs.host LIMIT ? OFFSET ?", limit, offset).Rows()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var host Host
		err = rows.Scan(&host.Id, &host.Uuid, &host.Name, &host.Hostname, &host.IpAddress, &host.Modified, &host.Severity)
		if err != nil {
			log.Print(err)
			return
		}
		i, err := strconv.ParseInt(host.Modified, 10, 64)
		if err != nil {
			panic(err)
		}
		host.Modified = time.Unix(i, 0).Format(time.RFC850)
		if host.Hostname.Valid == false {
			host.Hostname.String = ""
		}
		if host.IpAddress.Valid == false {
			host.IpAddress.String = ""
		}
		hosts = append(hosts, host)
	}

	var count int
	row := db.Raw("Select count(*) from hosts").Row() // (*sql.Row)
	row.Scan(&count)
	paginator := Paging(page, limit, offset, count, &hosts)
	json.NewEncoder(w).Encode(paginator)
}

func getHost(w http.ResponseWriter, r *http.Request) {
	// db, err := gorm.Open("postgres", "host=112.137.129.225 user=postgres dbname=gvmd password= sslmode=disable")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	vars := mux.Vars(r)
	var hosts []Host
	hostId := vars["id"]
	rows, err := db.Raw("select h.id, h.uuid, h.name, hostname.value hostname, ip.value ip, h.modification_time, hs.severity from hosts h left join (select distinct hi.host, hi.value from (select host, max(modification_time) as modification_time from host_identifiers where name='hostname' group by host) as hos inner join host_identifiers as hi on hi.host = hos.host and hi.modification_time = hos.modification_time where hi.name='hostname' and hi.value NOT LIKE 'ww%') AS hostname on h.id = hostname.host left join (select hi.host, hi.value from (select host, max(modification_time) as modification_time from host_identifiers where name='ip' group by host) as hos inner join host_identifiers as hi on hi.host = hos.host and hi.modification_time = hos.modification_time where hi.name='ip') AS ip on h.id = ip.host left join (select distinct hs.host, hs.severity from (select host, MAX(creation_time) as creation_time from host_max_severities group by host) as hms inner join host_max_severities hs on hs.host = hms.host and hs.creation_time = hms.creation_time) as hs ON h.id = hs.host WHERE h.uuid = ?", &hostId).Rows()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var host Host
		err = rows.Scan(&host.Id, &host.Uuid, &host.Name, &host.Hostname, &host.IpAddress, &host.Modified, &host.Severity)
		if err != nil {
			log.Print(err)
			return
		}
		i, err := strconv.ParseInt(host.Modified, 10, 64)
		if err != nil {
			panic(err)
		}
		// if host.Severity.Valid == false {
		// 	host.Severity.Float64= "N/A"
		// }
		host.Modified = time.Unix(i, 0).Format(time.RFC850)
		if host.Hostname.Valid == false {
			host.Hostname.String = ""
		}
		if host.IpAddress.Valid == false {
			host.IpAddress.String = ""
		}
		rowss, err := db.Raw("select i.uuid, i.name, i.value,i.source_id, i.source_data, i.creation_time from hosts h left join host_identifiers i on h.id =i.host where h.uuid = ? union all select ho.uuid, ho.name, oss.name as value,ho.source_id, ho.source_data, ho.creation_time from hosts h left join host_oss ho on h.id =ho.host inner join oss on ho.os = oss.id where h.uuid =  ? Order by creation_time DESC", &hostId, &hostId).Rows()
		if err != nil {
			log.Fatal(err)
		}
		var identifier Identifiers
		// var identifiers []Identifiers
		var source Sources
		var source_id string
		var source_name string
		for rowss.Next() {
			err = rowss.Scan(&identifier.Uuid, &identifier.Name, &identifier.Value, &source_id, &source_name, &identifier.Created)
			if err != nil {
				log.Print(err)
				return
			}
			i, err := strconv.ParseInt(identifier.Created, 10, 64)
			if err != nil {
				panic(err)
			}
			identifier.Created = time.Unix(i, 0).Format(time.RFC850)
			source.Report = source_id
			source.Nvt = source_name
			identifier.Source = source
			// identifiers = append(identifiers, identifier)
			host.Identifier = append(host.Identifier, identifier)
		}
		// host.Identifier = append(host.Identifier,identifiers)
		hosts = append(hosts, host)
	}
	json.NewEncoder(w).Encode(hosts)
}

func createHost(w http.ResponseWriter, r *http.Request) {
	var host Host
	_ = json.NewDecoder(r.Body).Decode(&host)
	var u = uuid.Must(uuid.NewV4())
	var created = time.Now().Unix()
	var modified = time.Now().Unix()

	db.Exec("INSERT INTO hosts(uuid, name, comment, creation_time, modification_time) VALUES (?,?,?,?,?)", &u, &host.Name, &host.Comment, &created, &modified)
	// if err != nil {
	// 	log.Print(err)
	// 	return
	// } else {
	// 	fmt.Fprintf(w, "Create Successful")
	// }

	// db.Create(&target)

}

func updateHost(w http.ResponseWriter, r *http.Request) {
	// db, err := gorm.Open("postgres", "host=112.137.129.225 user=postgres dbname=gvmd password= sslmode=disable")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	vars := mux.Vars(r)
	hostId := vars["id"]

	var host Host
	_ = json.NewDecoder(r.Body).Decode(&host)
	var modified = time.Now().Unix()

	db.Exec("UPDATE hosts SET name = ?, comment = ?, modification_time = ? WHERE id = ?", &host.Name, &host.Comment, &modified, &hostId)

}

func deleteHost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hostId := vars["id"]

	db.Exec("DELETE FROM hosts WHERE id = ?", &hostId)
}

func Paging(page int, limit int, offset int, count int, result interface{}) *Paginator {
	var paginator Paginator

	paginator.TotalRecord = count
	paginator.Records = result
	paginator.Offset = offset
	paginator.Page = page
	paginator.Limit = limit
	if page >= 1 {
		paginator.PrevPage = page - 1
	} else {
		paginator.PrevPage = page
	}

	if page == paginator.TotalPage {
		paginator.NextPage = page
	} else {
		paginator.NextPage = page + 1
	}

	paginator.TotalPage = int(math.Ceil(float64(paginator.TotalRecord) / float64(paginator.Limit)))
	return &paginator
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	fmt.Printf("%+v\n", user)
	fmt.Println(user.Password)
	// password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	// pass := string(password)
	rows, err := db.Raw("SELECT name, count(*) FROM Users WHERE name = ? AND password = ? GROUP BY name", &user.Name, &user.Password).Rows() // (*sql.Rows, error)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var count int
	for rows.Next() {
		rows.Scan(&user.Name, &count)
	}
	if count > 0 {
		// session.Values["id"] = &user.Id
		session.Values["username"] = &user.Name
		session.Save(r, w)
		fmt.Fprintf(w, "Login successful")

	} else {
		fmt.Fprintf(w, "Email or Password incorrect")
	}
}

func allUser(w http.ResponseWriter, r *http.Request) {
	// setupResponse(&w, r)
	// db, err := gorm.Open("postgres", "user=postgres dbname=mydb password=19121997 sslmode=disable")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	var users []User
	vars := mux.Vars(r)
	page, err := strconv.Atoi(vars["page"])

	var offset int
	limit := 10
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * limit
	}

	rows, err := db.Raw("SELECT u.id, u.uuid, u.name, ru.name, u.hosts, u.hosts_allow, u.ifaces_allow FROM users u INNER JOIN (select role_users.*, roles.name from role_users inner join roles on role_users.role = roles.id)as ru ON u.id = ru.user LIMIT ? OFFSET ?", limit, offset).Rows()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Role, &user.Hosts, &user.Host_allow, &user.Iface_allow)
		if err != nil {
			log.Print(err)
			return
		}
		if user.Hosts == "" {
			if user.Host_allow == "0" {
				user.Host_allow = "Allow all"
			} else if user.Host_allow == "1" {
				user.Host_allow = "Deny all"
			}
		} else {
			if user.Host_allow == "0" {
				user.Host_allow = "Allow all and deny from " + user.Hosts
			} else if user.Host_allow == "1" {
				user.Host_allow = "Deny all and allow from " + user.Hosts
			}
		}
		if user.Iface_allow == "0" || user.Iface_allow == "1" {
			user.Iface_allow = "Local"
		}

		users = append(users, user)
	}

	var count int
	row := db.Raw("Select count(*) from users").Row() // (*sql.Row)
	row.Scan(&count)
	paginator := Paging(page, limit, offset, count, &users)
	json.NewEncoder(w).Encode(paginator)
}
func createUser(w http.ResponseWriter, r *http.Request) {
	// db, err := gorm.Open("postgres", "user=postgres dbname=mydb password=19121997 sslmode=disable")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	// user = append(user, person)
	// json.NewEncoder(w).Encode(user)
	rows, err := db.Raw("SELECT count(*) FROM users WHERE name = ?", &user.Name).Rows() // (*sql.Rows, error)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var count int
	for rows.Next() {
		rows.Scan(&count)
	}
	if count > 0 {
		fmt.Fprintf(w, "User already exists")
	} else {
		var u = uuid.Must(uuid.NewV4())
		var created = time.Now().Unix()
		var modified = time.Now().Unix()
		// password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
		// pass := string(password)
		Host_allow, err := strconv.Atoi(user.Host_allow)
		if err == nil {
			fmt.Println("Error")
		}
		Iface_allow, err := strconv.Atoi(user.Iface_allow)
		if err == nil {
			fmt.Println("Error")
		}

		role, err := strconv.Atoi(user.Role)
		if err == nil {
			fmt.Println("Error")
		}
		db.Exec("INSERT INTO users(uuid, name, comment, password, hosts, hosts_allow, ifaces, ifaces_allow, creation_time, modification_time) VALUES (?,?,?,?,?,?,?,?,?,?)", &u, &user.Name, &user.Comment, &user.Password, &user.Hosts, &Host_allow, &user.Ifaces, &Iface_allow, &created, &modified)
		db.Exec("INSERT INTO role_users(role, user) VALUES(?,SELECT id FROM users WHERE name = ?)", &role, &user.Name)
	}

}

func updateUser(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// id := vars["id"]

	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	fmt.Fprintf(w, "Successfully Updated User")

}
func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	// db.Exec("DELETE FROM role_users WHERE user = ?", &userId)
	// db.Exec("DELETE FROM users WHERE id = ?", &userId)
	db.Exec("DELETE ur FROM dbo.UserRoles ur INNER JOIN dbo.Roles r ON r.RoleId = ur.Role INNER JOIN dbo.Users u ON ur.user = u.id WHERE u.id = ?", &userId)

}

// func handleRequests() {

// }
// func setupResponse(w *http.ResponseWriter, req *http.Request) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
// 	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// 	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
// }

func main() {
	// Handle Subsequent requests
	var err error
	db, err = gorm.Open("postgres", "host=112.137.129.225 user=postgres dbname=gvmd password= sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	initialMigration()
	// allowedHeaders := handlers.AllowedHeaders([]string{"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"})
	// allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	// allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/login", login).Methods("POST")
	myRouter.HandleFunc("/dashboard", getDashboard).Methods("GET")
	myRouter.HandleFunc("/users/page/{page}", allUser).Methods("GET")
	myRouter.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{id}", updateUser).Methods("PUT")
	myRouter.HandleFunc("/user", createUser).Methods("POST")
	myRouter.HandleFunc("/targets/page/{page}", allTargets).Methods("GET")
	myRouter.HandleFunc("/tasks/page/{page}", allTasks).Methods("GET")
	myRouter.HandleFunc("/target/{id}", getTarget).Methods("GET")
	myRouter.HandleFunc("/target", createTarget).Methods("POST")
	myRouter.HandleFunc("/target/{id}", updateTarget).Methods("PUT")
	myRouter.HandleFunc("/target/{id}", deleteTarget).Methods("DELETE")
	myRouter.HandleFunc("/nvts/page/{page}", allNvts).Methods("GET")
	myRouter.HandleFunc("/nvt/{id}", getNvt).Methods("GET")
	myRouter.HandleFunc("/cves/page/{page}", allCves).Methods("GET")
	myRouter.HandleFunc("/cve/{name}", getCve).Methods("GET")
	myRouter.HandleFunc("/cpes/page/{page}", allCpes).Methods("GET")
	myRouter.HandleFunc("/cpe/{id}", getCpe).Methods("GET")
	myRouter.HandleFunc("/hosts/page/{page}", allHosts).Methods("GET")
	myRouter.HandleFunc("/host/{id}", getHost).Methods("GET")
	myRouter.HandleFunc("/host", createHost).Methods("POST")
	myRouter.HandleFunc("/host/{id}", updateHost).Methods("PUT")
	myRouter.HandleFunc("/host/{id}", deleteHost).Methods("DELETE")
	myRouter.HandleFunc("/portlist", allPortList).Methods("GET")
	// corsObj := handlers.AllowedOrigins([]string{"*"})
	// http.ListenAndServe(":8081", handlers.CORS(corsObj)(myRouter))
	// http.ListenAndServe(":8081", handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(myRouter))
	log.Fatal(http.ListenAndServe(":8081", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "DELETE", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(myRouter)))
}
