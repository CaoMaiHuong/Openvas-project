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

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

var db *gorm.DB

type Dashboard struct {
	UserNumber string `json:"userNumber"`
	NvtNumber  string `json:"nvtNumber"`
	CveNumber  string `json:"cveNumber"`
	CpeNumber  string `json:"cpeNumber"`
}

type BySeverity struct {
	Total  int           `json:"total"`
	High   sql.NullInt64 `json:"high"`
	Medium sql.NullInt64 `json:"medium"`
	Low    sql.NullInt64 `json:"low"`
	Log    sql.NullInt64 `json:"log"`
	NA     sql.NullInt64 `json:"na"`
}

type Roles struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	Id                 int    `json:"id"`
	Uuid               string `json:"uuid"`
	Name               string `json:"name"`
	Owner              string `json:"owner"`
	Comment            string `json:"comment"`
	Role               string `json:"role"`
	RoleId             string `json:"role_id"`
	Password           string `json:"password"`
	Host_allow         string `json:"host_allow"`
	Host_allow_number  string `json:"host_allow_number"`
	Hosts              string `json:"hosts"`
	Iface_allow        string `json:"iface_allow"`
	Iface_allow_number string `json:"iface_allow_number"`
	Ifaces             string `json:"ifaces"`
	Created            string `json:"created"`
	Modified           string `json:"modified"`
	Message            string `json:"message"`
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
	PortListID    string        `json:"portlist_id"`
	Comment       string        `json:"comment"`
	MaxNumberHost string        `json:"maxhost"`
	RLOnly        string        `json:"rlonly"`
	RLUnify       string        `json:"rlunify"`
	Created       string        `json:"created"`
	Modified      string        `json:"modified"`
	AliveTest     string        `json:"alivetest"`
	Task          []Target_Task `json:"task"`
	Owner         string        `json:"owner"`
}

type Task struct {
	Id           int            `json:"id"`
	Uuid         string         `json:"uuid"`
	Name         string         `json:"name"`
	Owner        string         `json:"owner"`
	Status       string         `json:"status"`
	ReportNumber sql.NullString `json:"rpnumber"`
	// Reports      string         `json:"report"`
	LastReport             string         `json:"last_report"`
	Severity               sql.NullString `json:"severity"`
	Comment                string         `json:"comment"`
	Target                 string         `json:"target"`
	Alert                  string         `json:"alert"`
	Schedule               sql.NullString `json:"schedule"`
	In_assets              string         `json:"in_assets"`
	Assets_apply_overrides string         `json:"assets_apply_overrides"`
	Assets_min_qod         string         `json:"assets_min_qod"`
	Alterable              int            `json:"alterable"`
	Auto_delete            string         `json:"auto_delete"`
	Auto_delete_data       string         `json:"auto_delete_data"`
	Scanner                int            `json:"scanner"`
	Config                 int            `json:"config"`
	Network                string         `json:"network"`
	Hosts_ordering         string         `json:"hosts_ordering"`
	Max_checks             string         `json:"max_checks"`
	Max_hosts              string         `json:"max_hosts"`
}
type Task_Targert struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Scanner struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Config struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Report struct {
	Id       int            `json:"id"`
	Uuid     string         `json:"uuid"`
	Status   string         `json:"status"`
	Date     string         `json:"date"`
	Task     string         `json:"task"`
	Severity sql.NullString `json:"severity"`
	Rank     BySeverity     `json:"rank"`
}
type HostResult struct {
	Ip   string `json:"ip"`
	Name string `json:"name"`
}
type Result struct {
	Id            int        `json:"id"`
	Uuid          string     `json:"uuid"`
	Vulnerability string     `json:"vulnerability"`
	Severity      string     `json:"severity"`
	QoD           string     `json:"qod"`
	Host          HostResult `json:"host"`
	Location      string     `json:"location"`
	Created       string     `json:"created"`
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
	Id         int            `json:"id"`
	Uuid       string         `json:"uuid"`
	Name       string         `json:"name"`
	Comment    string         `json:"comment"`
	Hostname   sql.NullString `json:"hostname"`
	IpAddress  sql.NullString `json:"ipaddress"`
	Severity   sql.NullString `json:"severity"`
	Created    string         `json:"created"`
	Modified   string         `json:"modified"`
	Identifier []Identifiers  `json:"identifier"`
	Owner      string         `json:"owner"`
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
	var dashboards []Dashboard
	// rows, err := db.Raw("Select count(*) from users").Rows()
	// Scan(&dashboard.UserNumber)
	rows, err := db.Raw("Select count(*) from nvts").Rows()
	if err != nil {
		log.Print(err)
		return
	}
	for rows.Next() {
		err = rows.Scan(&dashboard.NvtNumber)
		if err != nil {
			log.Print(err)
			return
		}
	}
	rows1, err := db.Raw("Select count(*) from scap.cves").Rows()
	if err != nil {
		log.Print(err)
		return
	}
	for rows1.Next() {
		err = rows1.Scan(&dashboard.CveNumber)
		if err != nil {
			log.Print(err)
			return
		}

	}
	// Scan(&dashboard.NvtNumber)
	rows2, err := db.Raw("Select count(*) from scap.cpes").Rows()
	if err != nil {
		log.Print(err)
		return
	}
	for rows2.Next() {
		err = rows2.Scan(&dashboard.CpeNumber)
		if err != nil {
			log.Print(err)
			return
		}
	}
	// db.Raw("Select count(*) from scap.cpes").Scan(&dashboard.CpeNumber)
	// dashboards = append(dashboards, dashboard)
	dashboards = append(dashboards, dashboard)
	json.NewEncoder(w).Encode(dashboards)
}

func cveBySeverity(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Raw("SELECT count(*), sum(case when cvss between 7.0 and 10 then 1 else 0 end) as high, sum(case when  cvss between 4.0 and 6.9 then 1 else 0 end) as medium, sum(case when  cvss between 0.1 and 3.9 then 1 else 0 end) as low, sum(case when  cvss = 0 then 1 else 0 end) as log, sum(case when  cvss is null then 1 else 0 end) as na from scap.cves").Rows()
	if err != nil {
		log.Print(err)
		return
	}
	var cve BySeverity
	// var cves []CveBySeverity
	for rows.Next() {
		err = rows.Scan(&cve.Total, &cve.High, &cve.Medium, &cve.Low, &cve.Log, &cve.NA)
		if err != nil {
			log.Print(err)
			return
		}
	}
	// db.Raw("Select count(*) from scap.cpes").Scan(&dashboard.CpeNumber)
	// dashboards = append(dashboards, dashboard)
	// cves = append(cves, cve)
	json.NewEncoder(w).Encode(cve)
}
func nvtBySeverity(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Raw("SELECT count(*),sum(case when CAST(cvss_base as double precision) between 7.0 and 10 then 1 else 0 end) as high, sum(case when CAST(cvss_base as double precision) between 4.0 and 6.9 then 1 else 0 end) as medium, sum(case when CAST(cvss_base as double precision) between 0.1 and 3.9 then 1 else 0 end) as low, sum(case when CAST(cvss_base as double precision) = 0 then 1 else 0 end) as log, sum(case when CAST(cvss_base as double precision) is null then 1 else 0 end) as NA from nvts").Rows()
	if err != nil {
		log.Print(err)
		return
	}
	var nvt BySeverity
	// var cves []CveBySeverity
	for rows.Next() {
		err = rows.Scan(&nvt.Total, &nvt.High, &nvt.Medium, &nvt.Low, &nvt.Log, &nvt.NA)
		if err != nil {
			log.Print(err)
			return
		}
	}
	// db.Raw("Select count(*) from scap.cpes").Scan(&dashboard.CpeNumber)
	// dashboards = append(dashboards, dashboard)
	// cves = append(cves, cve)
	json.NewEncoder(w).Encode(nvt)
}
func allTargets(w http.ResponseWriter, r *http.Request) {

	// var paginator Paginator
	var targets []Target
	vars := mux.Vars(r)
	page, err := strconv.Atoi(vars["page"])

	var offset int
	limit := 6
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * limit
	}

	rows, err := db.Raw("SELECT t.id, t.uuid, t.name, t.hosts, t.reverse_lookup_only, t.reverse_lookup_unify, t.alive_test, p.id, p.name FROM targets t, port_lists p WHERE t.port_list = p.id LIMIT ? OFFSET ?", limit, offset).Rows()
	if err != nil {
		log.Print(err)
		return
	}

	for rows.Next() {
		var target Target
		err = rows.Scan(&target.Id, &target.Uuid, &target.Name, &target.Hosts, &target.RLOnly, &target.RLUnify, &target.AliveTest, &target.PortListID, &target.PortList)
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
	owner, err := strconv.Atoi(target.Owner)
	if err != nil {
		fmt.Println("Error")
	}
	var u = uuid.Must(uuid.NewV4())
	var created = time.Now().Unix()
	var modified = time.Now().Unix()
	fmt.Println(target)
	db.Exec("INSERT INTO targets(uuid,name,owner, hosts, reverse_lookup_only, reverse_lookup_unify, comment, port_list, alive_test, creation_time, modification_time) VALUES (?,?,(select id from users where id = ?),?,?,?,?,?,?,?,?)", &u, &target.Name, &owner, &target.Hosts, &target.RLOnly, &target.RLUnify, &target.Comment, &target.PortList, &target.AliveTest, &created, &modified)
	fmt.Fprintf(w, "Tạo target thành công!")
}

func updateTarget(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	targetId := vars["id"]

	var target Target
	_ = json.NewDecoder(r.Body).Decode(&target)
	var modified = time.Now().Unix()

	db.Exec("UPDATE targets SET name = ?, hosts = ?, reverse_lookup_only = ?, reverse_lookup_unify= ?, comment = ?, port_list = ?, alive_test = ?, modification_time = ? WHERE id = ?", &target.Name, &target.Hosts, &target.RLOnly, &target.RLUnify, &target.Comment, &target.PortListID, &target.AliveTest, &modified, &targetId)
	fmt.Fprintf(w, "Cập nhật thông tin thành công!")
}

func deleteTarget(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	targetId := vars["id"]

	db.Exec("INSERT INTO targets_trash(uuid, owner, name, hosts, exclude_hosts, reverse_lookup_only, reverse_lookup_unify, comment, port_list, alive_test, creation_time, modification_time) SELECT uuid, owner, name, hosts, exclude_hosts, reverse_lookup_only, reverse_lookup_unify, comment, port_list, alive_test, creation_time, modification_time FROM targets WHERE id = ?", &targetId)
	db.Exec("DELETE FROM targets WHERE id = ?", &targetId)
}

func allPortList(w http.ResponseWriter, r *http.Request) {
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

func allTaskTarget(w http.ResponseWriter, r *http.Request) {
	var targets []Target_Task
	rows, err := db.Raw("SELECT id, name from targets").Rows()
	if err != nil {
		log.Print(err)
		return
	}

	for rows.Next() {
		var target Target_Task
		err = rows.Scan(&target.Id, &target.Name)
		if err != nil {
			log.Print(err)
			return
		}
		targets = append(targets, target)
	}
	json.NewEncoder(w).Encode(targets)
}

func allScanner(w http.ResponseWriter, r *http.Request) {
	var scanners []Scanner
	rows, err := db.Raw("SELECT id, name from scanners").Rows()
	if err != nil {
		log.Print(err)
		return
	}

	for rows.Next() {
		var scanner Scanner
		err = rows.Scan(&scanner.Id, &scanner.Name)
		if err != nil {
			log.Print(err)
			return
		}
		scanners = append(scanners, scanner)
	}
	json.NewEncoder(w).Encode(scanners)
}

func allConfig(w http.ResponseWriter, r *http.Request) {
	var configs []Config
	rows, err := db.Raw("SELECT id, name from configs").Rows()
	if err != nil {
		log.Print(err)
		return
	}

	for rows.Next() {
		var config Config
		err = rows.Scan(&config.Id, &config.Name)
		if err != nil {
			log.Print(err)
			return
		}
		configs = append(configs, config)
	}
	json.NewEncoder(w).Encode(configs)
}

func allTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	vars := mux.Vars(r)
	page, err := strconv.Atoi(vars["page"])

	var offset int
	limit := 6
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * limit
	}

	rows, err := db.Raw("SELECT t.id, t.uuid, t.name, t.run_status, t.target, t.config, t.schedule, t.scanner, t.hosts_ordering, t.alterable, tp.max_checks, tp.max_hosts, tp.in_assets, tp.assets_apply_overrides, tp.assets_min_qod, tp.auto_delete, tp.auto_delete_data, r.count as reports, r.max as date, re.severity FROM tasks t LEFT JOIN (select task, count(id), max(date) from reports group by task) as r ON t.id = r.task LEFT JOIN (select task, max(severity) as severity from results group by task) as re ON t.id = re.task LEFT JOIN(select distinct tp.task, mc.value max_checks, mh.value max_hosts, ia.value in_assets, ao.value assets_apply_overrides, mq.value assets_min_qod, ad.value auto_delete, dd.value auto_delete_data from task_preferences tp inner join (select task, value from task_preferences where name='max_checks') as mc on tp.task = mc.task inner join (select task, value from task_preferences where name='max_hosts') as mh on tp.task = mh.task inner join (select task, value from task_preferences where name='in_assets') as ia on tp.task = ia.task inner join (select task, value from task_preferences where name='assets_apply_overrides') as ao on tp.task = ao.task inner join (select task, value from task_preferences where name='assets_min_qod') as mq on tp.task = mq.task inner join (select task, value from task_preferences where name='auto_delete') as ad on tp.task = ad.task inner join (select task, value from task_preferences where name='auto_delete_data') as dd on tp.task = dd.task) as tp on t.id = tp.task WHERE hidden = 0 LIMIT ? OFFSET ?", limit, offset).Rows()
	if err != nil {
		log.Print(err)
		return
	}

	for rows.Next() {
		var task Task
		var last sql.NullString
		err = rows.Scan(&task.Id, &task.Uuid, &task.Name, &task.Status, &task.Target, &task.Config, &task.Schedule, &task.Scanner, &task.Hosts_ordering, &task.Alterable, &task.Max_checks, &task.Max_hosts, &task.In_assets, &task.Assets_apply_overrides, &task.Assets_min_qod, &task.Auto_delete, &task.Auto_delete_data, &task.ReportNumber, &last, &task.Severity)
		if err != nil {
			log.Print(err)
			return
		}

		if last.Valid == false {
			last.String = ""
			task.LastReport = last.String
		} else {
			i, err := strconv.ParseInt(last.String, 10, 64)
			if err != nil {
				panic(err)
			}
			task.LastReport = time.Unix(i, 0).Format(time.RFC850)
		}
		switch task.Status {
		case "0":
			task.Status = "Delete Requested"
		case "1":
			task.Status = "Done"
		case "2":
			task.Status = "New"
		case "3":
			task.Status = "Requested"
		case "4":
			task.Status = "Running"
		case "10":
			task.Status = "Stop Requested"
		case "11":
			task.Status = "Ultimate Delete Requested"
		case "12":
			task.Status = "Stopped"
		case "13":
			task.Status = "Interrupted"
		case "14":
			task.Status = "Ultimate Delete Waiting"
		case "15":
			task.Status = "Stop Request Giveup"
		case "16":
			task.Status = "Deleted Waiting"
		case "17":
			task.Status = "Ultimate Delete Waiting"
		}

		tasks = append(tasks, task)
	}

	var count int
	row := db.Raw("Select count(*) from tasks where hidden = 0").Row() // (*sql.Row)
	row.Scan(&count)
	paginator := Paging(page, limit, offset, count, &tasks)
	json.NewEncoder(w).Encode(paginator)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	owner, err := strconv.Atoi(task.Owner)
	if err != nil {
		fmt.Println("Error")
	}
	var u = uuid.Must(uuid.NewV4())
	var created = time.Now().Unix()
	var modified = time.Now().Unix()

	db.Exec("with task as (insert into tasks(uuid, name, owner, comment, hidden, run_status, config, target, scanner, hosts_ordering, alterable, creation_time, modification_time) values (?,?,(select id from users where id=?),?,?,?,?,?,?,?,?,?,?) returning id) insert into task_preferences(task, name, value) values ((select id from task), 'max_checks', ?),((select id from task), 'max_hosts', ?),((select id from task), 'in_assets', ?),((select id from task), 'assets_apply_overrides', ?),((select id from task), 'assets_min_qod', ?),((select id from task), 'auto_delete', ?),((select id from task), 'auto_delete_data', ?)", &u, &task.Name, owner, &task.Comment, 0, 2, &task.Config, &task.Target, &task.Scanner, &task.Hosts_ordering, &task.Alterable, &created, &modified, &task.Max_checks, &task.Max_hosts, &task.In_assets, &task.Assets_apply_overrides, &task.Assets_min_qod, &task.Auto_delete, &task.Auto_delete_data)
	fmt.Fprintf(w, "Tạo tác vụ thành công!")
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId := vars["id"]

	var task Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	var modified = time.Now().Unix()
	db.Exec("with t as (update tasks set name=?, comment=?, config=?, target=?, scanner=?, hosts_ordering=?, alterable=?, modification_time=? where id = ? returning id) update task_preferences set value = (case when name = 'max_checks' then ? when name = 'max_hosts' then ? when name = 'in_assets' then ? when name = 'assets_apply_overrides' then ? when name = 'assets_min_qod' then ? when name = 'auto_delete' then ? when name = 'auto_delete_data' then ? end) from t WHERE task = t.id", &task.Name, &task.Comment, &task.Config, &task.Target, &task.Scanner, &task.Hosts_ordering, &task.Alterable, &modified, &taskId, &task.Max_checks, &task.Max_hosts, &task.In_assets, &task.Assets_apply_overrides, &task.Assets_min_qod, &task.Auto_delete, &task.Auto_delete_data)
	fmt.Fprintf(w, "Cập nhật thông tin thành công!")
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	// db, err := gorm.Open("postgres", "host=112.137.129.225 user=postgres dbname=gvmd password= sslmode=disable")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	vars := mux.Vars(r)
	taskId := vars["id"]

	db.Exec("UPDATE tasks SET hidden = 2 WHERE id = ?", &taskId)
}

func reportByTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId := vars["uuid"]
	page, err := strconv.Atoi(vars["page"])
	var reports []Report
	var offset int
	limit := 10
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * limit
	}
	rows, err := db.Raw("select r.id, r.uuid, r.date, r.scan_run_status, t.name from reports r inner join tasks t on r.task = t.id where t.uuid = ? LIMIT ? OFFSET ?", taskId, limit, offset).Rows()
	if err != nil {
		log.Print(err)
		return
	}

	for rows.Next() {
		var report Report
		err = rows.Scan(&report.Id, &report.Uuid, &report.Date, &report.Status, &report.Task)
		if err != nil {
			log.Print(err)
			return
		}
		switch report.Status {
		case "0":
			report.Status = "Delete Requested"
		case "1":
			report.Status = "Done"
		case "2":
			report.Status = "New"
		case "3":
			report.Status = "Requested"
		case "4":
			report.Status = "Running"
		case "10":
			report.Status = "Stop Requested"
		case "11":
			report.Status = "Ultimate Delete Requested"
		case "12":
			report.Status = "Stopped"
		case "13":
			report.Status = "Interrupted"
		case "14":
			report.Status = "Ultimate Delete Waiting"
		case "15":
			report.Status = "Stop Request Giveup"
		case "16":
			report.Status = "Deleted Waiting"
		case "17":
			report.Status = "Ultimate Delete Waiting"
		}
		i, err := strconv.ParseInt(report.Date, 10, 64)
		if err != nil {
			panic(err)
		}
		report.Date = time.Unix(i, 0).Format(time.RFC850)

		rows1, err := db.Raw("SELECT max(severity), sum(case when severity between 7.0 and 10 then 1 else 0 end) as high, sum(case when severity between 4.0 and 6.9 then 1 else 0 end) as medium, sum(case when severity between 0.1 and 3.9 then 1 else 0 end) as low, sum(case when severity = 0 then 1 else 0 end) as log, sum(case when severity is null then 1 else 0 end) as NA from results where report = ?", report.Id).Rows()
		if err != nil {
			log.Print(err)
			return
		}
		for rows1.Next() {
			var rp BySeverity
			err = rows1.Scan(&report.Severity, &rp.High, &rp.Medium, &rp.Low, &rp.Log, &rp.NA)
			if err != nil {
				log.Print(err)
				return
			}
			report.Rank = rp
			if report.Severity.Valid == true {
				ser, err := strconv.ParseFloat(report.Severity.String, 64)
				if err != nil {
					log.Print(err)
					return
				}
				if ser < 0 {
					report.Severity.String = "Error"
				}
			}
			if report.Severity.Valid == false {
				report.Severity.String = "N/A"
			}
		}

		reports = append(reports, report)
	}

	var count int
	row := db.Raw("select count(*) from reports r inner join tasks t on r.task = t.id where t.uuid = ?", taskId).Row() // (*sql.Row)
	row.Scan(&count)
	paginator := Paging(page, limit, offset, count, &reports)
	json.NewEncoder(w).Encode(paginator)
}

func getReport(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reportId := vars["uuid"]
	page, err := strconv.Atoi(vars["page"])
	var results []Result
	var ip string
	var name string
	var offset int
	limit := 6
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * limit
	}
	rows, err := db.Raw("select r.id, r.uuid, n.name, r.severity, r.qod, r.host, r.hostname, r.port, r.date from results r inner join nvts n on r.nvt = n.uuid inner join reports on r.report = reports.id where r.severity>=0 and reports.uuid = ? LIMIT ? OFFSET ?", reportId, limit, offset).Rows()
	if err != nil {
		log.Print(err)
		return
	}

	for rows.Next() {
		var result Result
		err = rows.Scan(&result.Id, &result.Uuid, &result.Vulnerability, &result.Severity, &result.QoD, &ip, &name, &result.Location, &result.Created)
		if err != nil {
			log.Print(err)
			return
		}
		i, err := strconv.ParseInt(result.Created, 10, 64)
		if err != nil {
			panic(err)
		}
		result.Created = time.Unix(i, 0).Format(time.RFC850)
		c := " %"
		// fmt.Println(s + string(c));
		result.QoD = result.QoD + string(c)
		var host HostResult
		host.Ip = ip
		host.Name = name
		result.Host = host
		results = append(results, result)
	}
	var count int
	row := db.Raw("select count(*) from results r inner join reports on r.report = reports.id where r.severity >=0 and reports.uuid = ?", reportId).Row() // (*sql.Row)
	row.Scan(&count)
	paginator := Paging(page, limit, offset, count, &results)
	json.NewEncoder(w).Encode(paginator)
}

func allNvts(w http.ResponseWriter, r *http.Request) {
	var nvts []Nvt
	vars := mux.Vars(r)
	page, err := strconv.Atoi(vars["page"])

	var offset int
	limit := 4
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

func searchNvts(w http.ResponseWriter, r *http.Request) {
	var nvts []Nvt
	vars := mux.Vars(r)
	search := vars["search"]
	page, err := strconv.Atoi(vars["page"])

	var offset int
	limit := 8
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * limit
	}

	rows, err := db.Raw("SELECT id, uuid, name, family, creation_time, modification_time, cve, cvss_base, qod FROM nvts WHERE Lower(name) LIKE '%'  || lower(?) || '%' LIMIT ? OFFSET ?", search, limit, offset).Rows()
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
		nvt.Qod = nvt.Qod + string(c)
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
		nvts = append(nvts, nvt)
	}

	var count int
	row := db.Raw("Select count(*) from nvts WHERE Lower(name) LIKE '%'  || lower(?) || '%'", search).Row() // (*sql.Row)
	row.Scan(&count)
	paginator := Paging(page, limit, offset, count, &nvts)
	json.NewEncoder(w).Encode(paginator)
}

func getNvt(w http.ResponseWriter, r *http.Request) {
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
		c := " %"
		nvt.Qod = nvt.Qod + string(c)
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
	var cves []Cve
	vars := mux.Vars(r)
	page, err := strconv.Atoi(vars["page"])

	var offset int
	limit := 3
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

func searchCves(w http.ResponseWriter, r *http.Request) {
	var cves []Cve
	vars := mux.Vars(r)
	search := vars["search"]
	page, err := strconv.Atoi(vars["page"])

	var offset int
	limit := 8
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * limit
	}

	rows, err := db.Raw("SELECT id, name, description, vector, complexity, authentication,  confidentiality_impact, integrity_impact, availability_impact, creation_time, cvss FROM scap.cves WHERE Lower(name) LIKE '%'  || lower(?) || '%' LIMIT ? OFFSET ?", search, limit, offset).Rows()
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
	row := db.Raw("Select count(*) from scap.cves WHERE Lower(name) LIKE '%'  || lower(?) || '%'", search).Row() // (*sql.Row)
	row.Scan(&count)
	paginator := Paging(page, limit, offset, count, &cves)
	json.NewEncoder(w).Encode(paginator)
}

func getCve(w http.ResponseWriter, r *http.Request) {
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

func searchCpes(w http.ResponseWriter, r *http.Request) {
	var cpes []Cpe
	vars := mux.Vars(r)
	search := vars["search"]
	page, err := strconv.Atoi(vars["page"])
	var offset int
	limit := 10
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * limit
	}
	rows, err := db.Raw("SELECT id, name, title, modification_time, cve_refs, max_cvss FROM scap.cpes WHERE Lower(title) LIKE '%'  || lower(?) || '%' LIMIT ? OFFSET ?", search, limit, offset).Rows()
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
	row := db.Raw("SELECT count(*) FROM scap.cpes WHERE Lower(title) LIKE '%'  || lower(?) || '%'", search).Row() // (*sql.Row)
	row.Scan(&count)
	paginator := Paging(page, limit, offset, count, &cpes)
	json.NewEncoder(w).Encode(paginator)
}
func allCpes(w http.ResponseWriter, r *http.Request) {
	var cpes []Cpe
	vars := mux.Vars(r)
	page, err := strconv.Atoi(vars["page"])
	var offset int
	limit := 6
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
	row := db.Raw("SELECT count(*) FROM scap.cpes").Row() // (*sql.Row)
	row.Scan(&count)
	paginator := Paging(page, limit, offset, count, &cpes)
	json.NewEncoder(w).Encode(paginator)
}
func getCpe(w http.ResponseWriter, r *http.Request) {
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
	var hosts []Host
	vars := mux.Vars(r)
	page, err := strconv.Atoi(vars["page"])

	var offset int
	limit := 6
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
		if host.Severity.Valid == false {
			host.Severity.String = "N/A"
		}
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
	vars := mux.Vars(r)
	var hosts []Host
	hostId := vars["id"]
	rows, err := db.Raw("select h.id, h.uuid, h.name, hostname.value hostname, ip.value ip,h.creation_time, h.modification_time, hs.severity from hosts h left join (select distinct hi.host, hi.value from (select host, max(modification_time) as modification_time from host_identifiers where name='hostname' group by host) as hos inner join host_identifiers as hi on hi.host = hos.host and hi.modification_time = hos.modification_time where hi.name='hostname' and hi.value NOT LIKE 'ww%') AS hostname on h.id = hostname.host left join (select hi.host, hi.value from (select host, max(modification_time) as modification_time from host_identifiers where name='ip' group by host) as hos inner join host_identifiers as hi on hi.host = hos.host and hi.modification_time = hos.modification_time where hi.name='ip') AS ip on h.id = ip.host left join (select distinct hs.host, hs.severity from (select host, MAX(creation_time) as creation_time from host_max_severities group by host) as hms inner join host_max_severities hs on hs.host = hms.host and hs.creation_time = hms.creation_time) as hs ON h.id = hs.host WHERE h.uuid = ?", &hostId).Rows()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var host Host
		err = rows.Scan(&host.Id, &host.Uuid, &host.Name, &host.Hostname, &host.IpAddress, &host.Created, &host.Modified, &host.Severity)
		if err != nil {
			log.Print(err)
			return
		}
		i, err := strconv.ParseInt(host.Modified, 10, 64)
		if err != nil {
			panic(err)
		}
		host.Modified = time.Unix(i, 0).Format(time.RFC850)
		j, err := strconv.ParseInt(host.Created, 10, 64)
		if err != nil {
			panic(err)
		}
		host.Created = time.Unix(j, 0).Format(time.RFC850)
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
	owner, err := strconv.Atoi(host.Owner)
	if err != nil {
		fmt.Println("Error")
	}
	var u = uuid.Must(uuid.NewV4())
	var created = time.Now().Unix()
	var modified = time.Now().Unix()
	rows, err := db.Raw("SELECT count(*) FROM hosts WHERE name = ?", host.Name).Rows() // (*sql.Rows, error)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var count int
	for rows.Next() {
		rows.Scan(&count)
	}
	if count > 0 {
		fmt.Fprintf(w, "Host đã tồn tại")
	} else {
		db.Exec("INSERT INTO hosts(uuid, name, owner, comment, creation_time, modification_time) VALUES (?,?,(select id from users where id = ?),?,?,?)", &u, &host.Name, &owner, &host.Comment, &created, &modified)
		fmt.Fprintf(w, "Tạo host thành công!")
	}
}

func updateHost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hostId := vars["id"]

	var host Host
	_ = json.NewDecoder(r.Body).Decode(&host)
	var modified = time.Now().Unix()

	db.Exec("UPDATE hosts SET name = ?, comment = ?, modification_time = ? WHERE id = ?", &host.Name, &host.Comment, &modified, &hostId)
	fmt.Fprintf(w, "Cập nhật thông tin thành công!")
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
	rows, err := db.Raw("SELECT id, name, count(*) FROM Users WHERE name = ? AND password = ? GROUP BY id,name", &user.Name, &user.Password).Rows() // (*sql.Rows, error)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var count int
	for rows.Next() {
		rows.Scan(&user.Id, &user.Name, &count)
	}
	if count > 0 {
		// session.Values["id"] = &user.Id
		session.Values["username"] = &user.Name
		session.Save(r, w)
		user.Message = "Login successful"

	} else {
		fmt.Fprintf(w, "Email or Password incorrect")
	}
	json.NewEncoder(w).Encode(user)
}

func allUser(w http.ResponseWriter, r *http.Request) {
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

	rows, err := db.Raw("SELECT u.id, u.uuid, u.name, u.password, ru.roleid, ru.name, u.hosts, u.hosts_allow, u.ifaces, u.ifaces_allow FROM users u INNER JOIN (select role_users.*, roles.id roleid, roles.name from role_users inner join roles on role_users.role = roles.id)as ru ON u.id = ru.user LIMIT ? OFFSET ?", limit, offset).Rows()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Password, &user.RoleId, &user.Role, &user.Hosts, &user.Host_allow_number, &user.Ifaces, &user.Iface_allow_number)
		if err != nil {
			log.Print(err)
			return
		}
		if user.Hosts == "" {
			if user.Host_allow_number == "0" {
				user.Host_allow = "Allow all"
			} else if user.Host_allow_number == "1" {
				user.Host_allow = "Deny all"
			}
		} else {
			if user.Host_allow_number == "0" {
				user.Host_allow = "Allow all and deny from " + user.Hosts
			} else if user.Host_allow_number == "1" {
				user.Host_allow = "Deny all and allow from " + user.Hosts
			}
		}
		if user.Iface_allow_number == "0" || user.Iface_allow_number == "1" {
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

func getRoles(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Raw("SELECT id, name FROM roles").Rows()
	if err != nil {
		log.Fatal(err)
	}
	var roles []Roles
	for rows.Next() {
		var role Roles
		err = rows.Scan(&role.Id, &role.Name)
		if err != nil {
			log.Print(err)
			return
		}
		roles = append(roles, role)
	}
	json.NewEncoder(w).Encode(roles)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	owner, err := strconv.Atoi(user.Owner)
	if err != nil {
		fmt.Println("Error")
	}
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
		fmt.Fprintf(w, "Người dùng đã tồn tại")
	} else {
		var u = uuid.Must(uuid.NewV4())
		var created = time.Now().Unix()
		var modified = time.Now().Unix()
		// password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
		// pass := string(password)
		Host_allow, err := strconv.Atoi(user.Host_allow_number)
		if err != nil {
			fmt.Println("Error")
		}
		Iface_allow, err := strconv.Atoi(user.Iface_allow_number)
		if err != nil {
			fmt.Println("Error")
		}
		role, err := strconv.Atoi(user.RoleId)
		if err != nil {
			fmt.Println("Error")
		}
		db.Exec("with userr as (insert into users(uuid, name, owner, comment, password, hosts, hosts_allow, ifaces, ifaces_allow, creation_time, modification_time) values (?,?,?,?,?,?,?,?,?,?,?) returning id) insert into role_users(role, \"user\") values ( ?, (select id from userr))", &u, &user.Name, owner, &user.Comment, &user.Password, &user.Hosts, &Host_allow, &user.Ifaces, &Iface_allow, &created, &modified, &role)
		fmt.Fprintf(w, "Tạo người dùng thành công!")
	}
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	var modified = time.Now().Unix()
	Host_allow, err := strconv.Atoi(user.Host_allow_number)
	if err != nil {
		fmt.Println("Error")
	}
	Iface_allow, err := strconv.Atoi(user.Iface_allow_number)
	if err != nil {
		fmt.Println("Error")
	}
	role, err := strconv.Atoi(user.RoleId)
	if err != nil {
		fmt.Println("Error")
	}
	db.Exec("with userr as (update users set name= ?, comment = ?, hosts = ?, hosts_allow = ?, ifaces = ?, ifaces_allow = ?, modification_time = ? where id = ? returning id) update role_users set role = ? from userr where \"user\" = userr.id", &user.Name, &user.Comment, &user.Hosts, &Host_allow, &user.Ifaces, &Iface_allow, &modified, &userId, &role)
	fmt.Fprintf(w, "Cập nhật thông tin thành công!")
}
func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	db.Exec("DELETE FROM role_users WHERE \"user\" = ?", &userId)
	db.Exec("DELETE FROM users WHERE id = ?", &userId)
}

func main() {
	// Handle Subsequent requests
	var err error
	db, err = gorm.Open("postgres", "host=112.137.129.225 user=postgres dbname=gvmd password= sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	// initialMigration()
	// allowedHeaders := handlers.AllowedHeaders([]string{"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"})
	// allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	// allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/login", login).Methods("POST")

	// Dashboard
	myRouter.HandleFunc("/dashboard", getDashboard).Methods("GET")
	myRouter.HandleFunc("/cvebyseverity", cveBySeverity).Methods("GET")
	myRouter.HandleFunc("/nvtbyseverity", nvtBySeverity).Methods("GET")

	// Users
	myRouter.HandleFunc("/users/page/{page}", allUser).Methods("GET")
	myRouter.HandleFunc("/roles", getRoles).Methods("GET")
	myRouter.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{id}", updateUser).Methods("PUT")
	myRouter.HandleFunc("/user", createUser).Methods("POST")

	// Tasks
	myRouter.HandleFunc("/tasks/page/{page}", allTasks).Methods("GET")
	myRouter.HandleFunc("/task", createTask).Methods("POST")
	myRouter.HandleFunc("/task/{id}", updateTask).Methods("PUT")
	myRouter.HandleFunc("/task/{id}", deleteTask).Methods("DELETE")
	myRouter.HandleFunc("/task_target", allTaskTarget).Methods("GET")
	myRouter.HandleFunc("/scanners", allScanner).Methods("GET")
	myRouter.HandleFunc("/configs", allConfig).Methods("GET")
	myRouter.HandleFunc("/reports/{uuid}/page/{page}", reportByTask).Methods("GET")
	myRouter.HandleFunc("/report/{uuid}/page/{page}", getReport).Methods("GET")

	// Targets
	myRouter.HandleFunc("/targets/page/{page}", allTargets).Methods("GET")
	myRouter.HandleFunc("/target/{id}", getTarget).Methods("GET")
	myRouter.HandleFunc("/target", createTarget).Methods("POST")
	myRouter.HandleFunc("/target/{id}", updateTarget).Methods("PUT")
	myRouter.HandleFunc("/target/{id}", deleteTarget).Methods("DELETE")

	//Nvts
	myRouter.HandleFunc("/nvts/page/{page}", allNvts).Methods("GET")
	myRouter.HandleFunc("/nvt/{id}", getNvt).Methods("GET")
	myRouter.HandleFunc("/nvts/page/{page}/search/{search}", searchNvts).Methods("GET")

	// Cves
	myRouter.HandleFunc("/cves/page/{page}", allCves).Methods("GET")
	myRouter.HandleFunc("/cve/{name}", getCve).Methods("GET")
	myRouter.HandleFunc("/cves/page/{page}/search/{search}", searchCves).Methods("GET")

	// Cpes
	myRouter.HandleFunc("/cpes/page/{page}", allCpes).Methods("GET")
	myRouter.HandleFunc("/cpe/{id}", getCpe).Methods("GET")
	myRouter.HandleFunc("/cpes/page/{page}/search/{search}", searchCpes).Methods("GET")

	// Hosts
	myRouter.HandleFunc("/hosts/page/{page}", allHosts).Methods("GET")
	myRouter.HandleFunc("/host/{id}", getHost).Methods("GET")
	myRouter.HandleFunc("/host", createHost).Methods("POST")
	myRouter.HandleFunc("/host/{id}", updateHost).Methods("PUT")
	myRouter.HandleFunc("/host/{id}", deleteHost).Methods("DELETE")
	myRouter.HandleFunc("/portlist", allPortList).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "DELETE", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(myRouter)))
}
