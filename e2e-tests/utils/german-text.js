/**
 * German text constants for assertions in E2E tests
 * All UI text is in German, so we need these for reliable assertions
 */

const GERMAN_TEXT = {
  // Authentication
  LOGIN_SUCCESS: 'Erfolgreich angemeldet',
  LOGIN_ERROR: 'Ungültige Anmeldedaten',
  REGISTER_SUCCESS: 'Registrierung erfolgreich',
  LOGOUT_SUCCESS: 'Erfolgreich abgemeldet',
  EMAIL_VERIFICATION_SENT: 'Bestätigungs-E-Mail',
  PASSWORD_RESET_SENT: 'E-Mail wurde gesendet',
  PASSWORD_CHANGED: 'Passwort erfolgreich geändert',

  // Validation errors
  INVALID_EMAIL: 'Ungültige E-Mail',
  INVALID_PASSWORD: 'Passwort',
  REQUIRED_FIELD: 'Pflichtfeld',
  PASSWORD_MISMATCH: 'Passwörter stimmen nicht überein',
  TERMS_REQUIRED: 'Nutzungsbedingungen',

  // Bookings
  BOOKING_CREATED: 'Buchung erfolgreich',
  BOOKING_CANCELLED: 'storniert',
  BOOKING_MOVED: 'verschoben',
  CANNOT_CANCEL: 'nicht storniert werden',
  DOUBLE_BOOKING_ERROR: 'bereits gebucht',
  BLOCKED_DATE_ERROR: 'gesperrt',
  EXPERIENCE_LEVEL_ERROR: 'Erfahrungsstufe',
  BOOKING_NOTES_ADDED: 'Notiz',

  // Profile
  PROFILE_UPDATED: 'Profil aktualisiert',
  PHOTO_UPLOADED: 'Foto hochgeladen',
  ACCOUNT_DELETED: 'Konto gelöscht',

  // Dogs
  DOG_CREATED: 'Hund erstellt',
  DOG_UPDATED: 'Hund aktualisiert',
  DOG_DELETED: 'Hund gelöscht',
  DOG_UNAVAILABLE: 'nicht verfügbar',

  // Admin
  USER_DEACTIVATED: 'deaktiviert',
  USER_ACTIVATED: 'aktiviert',
  REQUEST_APPROVED: 'genehmigt',
  REQUEST_DENIED: 'abgelehnt',
  SETTINGS_UPDATED: 'Einstellungen gespeichert',

  // General
  SUCCESS: 'Erfolgreich',
  ERROR: 'Fehler',
  CONFIRM: 'Bestätigen',
  CANCEL: 'Abbrechen',
  SAVE: 'Speichern',
  DELETE: 'Löschen',
  EDIT: 'Bearbeiten',

  // Navigation
  HOME: 'Startseite',
  LOGIN: 'Anmelden',
  REGISTER: 'Registrieren',
  DASHBOARD: 'Dashboard',
  PROFILE: 'Profil',
  DOGS: 'Hunde',
  CALENDAR: 'Kalender',
  LOGOUT: 'Abmelden',
};

module.exports = GERMAN_TEXT;

// DONE: German text constants for reliable test assertions
