import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthTokenService } from '../auth-token.service';


@Component({
  selector: 'app-manager',
  templateUrl: './manager.component.html',
  styleUrls: ['./manager.component.css']
})
export class ManagerComponent implements OnInit {


  eid: string | null;
  constructor(private authTokenService : AuthTokenService,
    private router: Router) { }

  ngOnInit(): void {
    this.eid = this.authTokenService.getEmployeeID()
  }

  viewEmployees(){
    this.router.navigate(['employee-list',this.eid])
  }

  addEmployee(){
    this.router.navigate(['create-employee',this.eid])
  }

  viewPersonalInfo(){
    this.router.navigate(['personal',this.eid])
  }
}
