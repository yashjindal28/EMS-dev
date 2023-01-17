import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { AuthTokenService } from '../auth-token.service';
import { Employee } from '../employee';
import { EmployeeService } from '../employee.service';

@Component({
  selector: 'app-manager-update-employee',
  templateUrl: './manager-update-employee.component.html',
  styleUrls: ['./manager-update-employee.component.css']
})
export class ManagerUpdateEmployeeComponent implements OnInit {

  eid : string | null;
  childEid: string;
  managerID : string;
  employee : Employee = new Employee();
  constructor(private employeeService:EmployeeService,
    private route: ActivatedRoute, private router:Router,
    private authTokenService: AuthTokenService) { }

  ngOnInit(): void {
    this.eid = this.authTokenService.getEmployeeID()
    this.childEid = this.route.snapshot.params['childEid'];
    this.employeeService.getEmployeeById(this.eid,this.childEid).subscribe(data => {
      this.employee = data;
      this.managerID = this.employee.managerID
    }, error => console.log(error));
  }

  onSubmit(){
    this.employeeService.updateEmployee(this.eid,this.childEid,this.employee).subscribe(data =>{
      //this.goToEmployeeList();
    }, 
     error => console.log(error))

     this.employeeService.updateEmployeePayroll(this.eid,this.childEid,this.employee).subscribe(data =>{
      //this.goToEmployeeList();
    }, 
     error => console.log(error))

     this.employeeService.updateEmployeeLeave(this.eid,this.childEid,this.employee).subscribe(data =>{
      //this.goToEmployeeList();
    }, 
     error => console.log(error))

     this.employeeService.updateEmployeeDepartment(this.eid,this.childEid,this.employee).subscribe(data =>{
      //this.goToEmployeeList();
    }, 
     error => console.log(error))
    

     this.goToEmployeeList();
  }

  goToEmployeeList(){
    this.router.navigate(['/employee-list',this.managerID]);
  }
  
}
