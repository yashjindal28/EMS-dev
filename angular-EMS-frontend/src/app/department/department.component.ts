import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthTokenService } from '../auth-token.service';
import { Departments } from '../employee';
import { EmployeeService } from '../employee.service';

@Component({
  selector: 'app-department',
  templateUrl: './department.component.html',
  styleUrls: ['./department.component.css']
})
export class DepartmentComponent implements OnInit {

  departments: Departments[]
  eid: string | null
  constructor(private employeeService: EmployeeService,
    private router:Router, private authTokenService: AuthTokenService) { }

  ngOnInit(): void {
    this.eid = this.authTokenService.getEmployeeID()
    this.employeeService.getAllDepartments(this.eid).subscribe(data => {
      this.departments = data;
    });
  }

  viewProjects(did: string){
    this.router.navigate(['department-projects',did])
  }

  viewEmployees(did: string){
    this.router.navigate(['department-employees',did])
  }
}
