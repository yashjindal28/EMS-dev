import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { AuthTokenService } from '../auth-token.service';
import { Employee } from '../employee';
import { EmployeeService } from '../employee.service';

@Component({
  selector: 'app-manager-employee-detail',
  templateUrl: './manager-employee-detail.component.html',
  styleUrls: ['./manager-employee-detail.component.css']
})
export class ManagerEmployeeDetailComponent implements OnInit {

  eid: string | null;
  childEid: string;
  employee: Employee
  constructor(private route:ActivatedRoute,
    private employeeService: EmployeeService,
    private authTokenService: AuthTokenService) { }

  ngOnInit(): void {
    this.eid = this.authTokenService.getEmployeeID()
    this.childEid = this.route.snapshot.params['childEid']

    this.employee = new Employee();
    this.employeeService.getEmployeeById(this.eid,this.childEid).subscribe( data => {
      this.employee = data;
    });
  }

}
