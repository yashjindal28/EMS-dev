import { Component } from '@angular/core';
import { AuthTokenService } from './auth-token.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'angular-EMS-frontend';

  eid: string | null;
  constructor(public authTokenService: AuthTokenService) {
    this.eid = this.authTokenService.getEmployeeID()
   }

  public isLoggedIn(){
    this.eid = this.authTokenService.getEmployeeID()
    return this.authTokenService.isLoggedIn()
  }

  public logout(){
    this.authTokenService.clear()
  }

}
