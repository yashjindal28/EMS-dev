import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class AuthTokenService {

  constructor() { }

  public setToken(jwtToken:string) {
    localStorage.setItem("jwtToken",jwtToken)
  }

  public getToken(): string | null{
    return localStorage.getItem('jwtToken')
  }

  public setDesg(desg:string) {
    localStorage.setItem("desg",desg)
  }

  public getDesg(): string | null{
    return localStorage.getItem('desg')
  }

  public setEmployeeID(eid:string) {
    localStorage.setItem("eid",eid)
  }

  public getEmployeeID(): string | null{
    return localStorage.getItem('eid')
  }

  public clear(){
    localStorage.clear()
  }

  public isLoggedIn() {
    if(this.getToken == null || this.getDesg() == null || this.getEmployeeID()== null)
          return false
    else
      return true
  }

  public roleMatch(allowedRole: string): boolean{
    if(allowedRole == this.getDesg()){
      return true
    } else {
      return false
    }
  }

  public setWrong(){
    localStorage.setItem("wrongCred","true")
  }

  public getWrong(): string | null{
    return localStorage.getItem('wrongCred')
  }
}
