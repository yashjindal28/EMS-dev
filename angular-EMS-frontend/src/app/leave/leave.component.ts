import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthTokenService } from '../auth-token.service';
import { Employee } from '../employee';

@Component({
  selector: 'app-leave',
  templateUrl: './leave.component.html',
  styleUrls: ['./leave.component.css']
})
export class LeaveComponent implements OnInit {

  employee: Employee = new Employee();
  constructor(private router: Router,
    private authTokenService: AuthTokenService) { }

  ngOnInit(): void {

    this.employee.eid = String(this.authTokenService.getEmployeeID())

  }

  callManager(eid: string) {
    this.router.navigate(['leave-manager', eid])
  }

  callEmployee(eid: string) {
    this.router.navigate(['leave-employee', eid])
  }

}
