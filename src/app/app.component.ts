import { AuthService } from '@app/auth/auth.service';
import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { take } from 'rxjs/operators';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {

  constructor(public authService: AuthService, private httpClient: HttpClient) {
    authService.handleAuthentication();
  }

  getStatus() {
    this.httpClient.get<any>('http://localhost:3000/api/v1/version').pipe(take(1)).subscribe(result => alert(JSON.stringify(result)));
  }
}
