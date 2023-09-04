import { Injectable } from '@angular/core';
import { KeycloakService } from 'keycloak-angular';
import { Observable, from, map } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  isLoggedIn$: Observable<boolean>;
  isLoggedOut$: Observable<boolean>;

  constructor(private keycloak: KeycloakService) {
    this.isLoggedIn$ = from(this.keycloak.isLoggedIn());
    this.isLoggedOut$ = this.isLoggedIn$.pipe(map((loggedIn) => !loggedIn));
  }

  async logout() {
    await this.keycloak.logout(window.location.origin);
  }
}
