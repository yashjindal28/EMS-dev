import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Employee } from '../employee';
import { EmployeeService } from '../employee.service';

@Component({
  selector: 'app-manager-employee-list',
  templateUrl: './manager-employee-list.component.html',
  styleUrls: ['./manager-employee-list.component.css']
})
export class ManagerEmployeeListComponent implements OnInit {

  employees: Employee[];
  eid : string;
  constructor(private employeeService: EmployeeService,
    private router: Router, private route: ActivatedRoute) { }

  ngOnInit(): void {
    this.getEmployees();
  }

  private getEmployees(){
    this.eid = this.route.snapshot.params['eid']
    this.employeeService.getEmployeesUnderManager(this.eid).subscribe(data => {
      this.employees = data;
    
    });
  }

  updateEmployee(childEid: string){
    this.router.navigate(['update-employee',childEid]);
  }

  deleteEmployee(childEid: string,did: string){
    this.employeeService.deleteEmployee(this.eid,childEid).subscribe(data => {
     
    });

    this.employeeService.deleteEmployeeFromDepartment(this.eid,childEid,did).subscribe(data => {
     
    });

    this.employeeService.deleteEmployeeFromLeave(this.eid,childEid).subscribe(data => {
     
    });

    this.employeeService.deleteEmployeeFromPayroll(this.eid,childEid).subscribe(data => {
     
    });

    this.employeeService.deleteUser(this.eid,childEid).subscribe(data => {
     
    });

    setTimeout(() => {
    }, 500);
    this.getEmployees()
  }

  viewEmployee(childEid : string){
    this.router.navigate(['employee-detail',childEid]);
  }
}
