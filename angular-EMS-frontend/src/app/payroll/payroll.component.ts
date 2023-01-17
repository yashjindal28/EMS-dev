import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthTokenService } from '../auth-token.service';
import { Employee } from '../employee';
import { EmployeeService } from '../employee.service';

@Component({
  selector: 'app-payroll',
  templateUrl: './payroll.component.html',
  styleUrls: ['./payroll.component.css']
})
export class PayrollComponent implements OnInit {

  employees : Employee[];
  eid: string | null;
  constructor(private employeeService: EmployeeService,
    private router:Router,private authTokenService: AuthTokenService) { }

  ngOnInit(): void {

    this.eid = this.authTokenService.getEmployeeID()

    this.employeeService.getAllEmployeesForPayroll(this.eid).subscribe(data => {
      this.employees = data;    
    });

  }

  issuePayCheck(childEid : string){
    this.employeeService.issuePayCheck(this.eid,childEid).subscribe(data => {
         
    });
    this.router.routeReuseStrategy.shouldReuseRoute = () => false;
    this.router.onSameUrlNavigation = 'reload';
    this.router.navigate(['payroll'])
  }

}
