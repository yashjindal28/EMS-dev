import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Count, Employee } from '../employee';
import { EmployeeService } from '../employee.service';

@Component({
  selector: 'app-manager-create-employee',
  templateUrl: './manager-create-employee.component.html',
  styleUrls: ['./manager-create-employee.component.css']
})
export class ManagerCreateEmployeeComponent implements OnInit {

  employee : Employee = new Employee();
  manager: Employee = new Employee();
  count : Count = new Count();
  countDpt: number;
  countOfEmp: number;
  managerID : string;
  
  prefix: string = "EXE";
  resultforEXE : boolean
  constructor(public employeeService: EmployeeService,
    private router: Router, private route: ActivatedRoute) { }

  ngOnInit(): void {
    this.employee.managerID = this.route.snapshot.params['eid'];
    this.resultforEXE = this.employee.managerID.startsWith(this.prefix)

    this.employeeService.getEmployeeById(this.employee.managerID,this.employee.managerID).subscribe(data => {
      this.manager = data;
      this.employee.manager = this.manager.firstname + " " + this.manager.lastname;
      this.employee.dpt = this.manager.dpt;
      this.employee.did = this.manager.did;
      
    }, error => console.log(error));

    setTimeout(() => {
    }, 500);

    this.employeeService.getCount(this.employee.managerID).subscribe(data => {
      this.count = data;

      this.countDpt = this.count.countOfDpt
      this.countOfEmp = this.count.countOfEmp

      if (this.employee.managerID == "admin"){
        this.employee.dpt = "Executive"
        this.employee.did = "D1"
        this.employee.manager = "admin"
        this.employee.eid = "EXE1"
        this.employee.desg = "CEO"
      } else if (this.resultforEXE == true){
        this.employee.did = "D" + String(this.countDpt+1);
      } else {
        this.employee.eid = this.employee.dpt.substring(0,3) + String(this.count.countOfEmp+1);
      }

    }, error => console.log(error));

    

   
    
    //this.employee.eid = this.employee.dpt + "";
  }

  saveEmployee(){
    if (this.employee.managerID == "EXE1"){
      this.employee.dpt = (document.getElementById('dpt') as HTMLInputElement).value
    }
    this.employeeService.createEmployee(this.employee.managerID,this.employee).subscribe( data =>{
      console.log(data);
    },
    error => {console.log(error); return });

    this.employeeService.createUser(this.employee.managerID,this.employee).subscribe( data =>{
    
    },
    error => {console.log(error); return });

    this.employeeService.createPayrollInfoForNewEmployee(this.employee.managerID,this.employee).subscribe( data =>{

    },
    error => {console.log(error); return });

    this.employeeService.createLeaveInfoForNewEmployee(this.employee.managerID,this.employee).subscribe( data =>{

    },
    error => {console.log(error); return });

    this.employeeService.addNewEmployeeToDepartment(this.employee.managerID,this.employee).subscribe( data =>{

    },
    error => {console.log(error); return });

    this.goToEmployeeList();
  }

  goToEmployeeList(){
    this.router.navigate(['/employee-list',this.employee.managerID]);
  }

  onSubmit(){
    console.log(this.employee);
    this.saveEmployee();
  }
  
  public text: string;
  onChange(updatedValue : string){
    
    console.log(updatedValue)
    if (updatedValue.length < 4){
      this.employee.eid = updatedValue.toUpperCase() + String(this.count.countOfEmp+1);
    }
    
  }
}
