import smtplib
from email.message import EmailMessage

SMTP_HOST = "127.0.0.1"
SMTP_PORT = 2525
SENDER_EMAIL = "test-sender@local.dev"
RECIPIENT_EMAIL = "test-recipient@local.dev"


def main() -> None:
    msg = EmailMessage()
    msg["From"] = SENDER_EMAIL
    msg["To"] = RECIPIENT_EMAIL
    msg["Subject"] = "Mailpit test"
    msg.set_content("MAILPIt TEST")

    with smtplib.SMTP(SMTP_HOST, SMTP_PORT, timeout=10) as server:
        server.send_message(msg)

    print(f"Письмо отправлено через {SMTP_HOST}:{SMTP_PORT}")


if __name__ == "__main__":
    main()
