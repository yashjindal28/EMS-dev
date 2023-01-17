import { Component, OnInit } from '@angular/core';
import { AuthTokenService } from '../auth-token.service';
import { Employee } from '../employee';
import { EmployeeService } from '../employee.service';

@Component({
  selector: 'app-allemployees',
  templateUrl: './allemployees.component.html',
  styleUrls: ['./allemployees.component.css']
})
export class AllemployeesComponent implements OnInit {
  
  employees : Employee[];
  eid: string|null;
  enteredSearchValue : string = ''
  criteriaFilter : Employee = new Employee()
  constructor(private employeeService: EmployeeService,
    private authTokenService: AuthTokenService) { }

  ngOnInit(): void {
    this.eid = this.authTokenService.getEmployeeID()
    this.employeeService.getAllEmployees(this.eid).subscribe(data => {
      this.employees = data;    
    });
  }

  onSubmit(){

    this.employeeService.searchAllEmployees(this.eid,this.enteredSearchValue).subscribe(data => {
      this.employees = data;

    });
  }

  filter(){
    this.employeeService.filterAllEmployees(this.eid,this.criteriaFilter).subscribe(data => {
      this.employees = data;

    });
  }

}
