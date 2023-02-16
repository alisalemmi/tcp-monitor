# TCP Monitoring

- [EN](#-en)
- [فا](#-فارسی)

## 🌎 EN

This app is a prometheus exporter written in golang that can be
used to monitor the retransmission of TCP packets.
You can also use the
[docker image](https://hub.docker.com/r/alisalemmi/tcp-monitor)
of this app. To use this image you should create a volume that
mount `/proc/net` to `/host/net` and also **enable host networking**.

## 🌍 فارسی

این برنامه ابزاری برای نظارت بر نرخ باز ارسال بسته‌های
TCP
است که به زبان
golang
نوشته شده است. از این برنامه می‌توان به عنوان
exporter
برای
prometheus
استفاده کرد. برای اجرای این برنامه می‌توان از
image
داکر آن که در
[این لینک](https://hub.docker.com/r/alisalemmi/tcp-monitor)
قرار دارد استفاده کرد.
برای عمل کرد صحیح این برنامه لازم است مسیر
&lrm;`/proc/net`&rlm;
در ماشین میزبان به مسیر
&lrm;`/host/net`&rlm;
در کانتینر
mount
شود. علاوه بر این لازم است حالت شبکه روی میزبان تنظیم شود.
برای این کار می‌توان هنگام ساخت کانتینر از آرگومان
&lrm;`--net host`&rlm;
استفاده کرد.
