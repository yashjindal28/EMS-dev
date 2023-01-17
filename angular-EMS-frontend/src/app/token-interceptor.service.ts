import { HttpEvent, HttpHandler, HttpInterceptor, HttpRequest } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { AuthTokenService } from './auth-token.service';

@Injectable({
  providedIn: 'root'
})
export class TokenInterceptorService implements HttpInterceptor {

  constructor(private authTokenService: AuthTokenService) { }
  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    
    let token = this.authTokenService.getToken()
    let jwtToken = req.clone({
      setHeaders:{
        Authorization: 'Bearer ' + token
      }
    })
    return next.handle(jwtToken)
  }
}
