import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthTokenService } from '../auth-token.service';
import { AuthService } from '../auth.service';
import { User } from '../employee';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  user: User = new User()
  wrongCred: boolean
  constructor(private authService: AuthService,
    private router: Router, private authTokenService: AuthTokenService) { }

  ngOnInit(): void {
    if (this.authTokenService.getWrong() === "true")
        this.wrongCred = true
    else
        this.wrongCred = false
  }

  login() {
    this.authService.login(this.user).subscribe((data: any) => {
      this.authTokenService.setToken(data.jwtToken)
      this.authTokenService.setDesg(data.desg)
      this.authTokenService.setEmployeeID(this.user.eid)

      const desg = data.desg
      if (desg === "Manager" || desg === "CEO") {
        this.router.navigate(['manager'])
      } else {
        this.router.navigate(['employee'])
      }

    }, error => this.wrong())

  }

  wrong(){
    this.authTokenService.setWrong()
    this.router.routeReuseStrategy.shouldReuseRoute = () => false;
    this.router.onSameUrlNavigation = 'reload';
    this.router.navigate(['login'])
  }

}
