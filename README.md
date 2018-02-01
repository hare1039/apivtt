# apivtt
An api server that accept subtitles from url and returns webvtt-encoded subtitle `.vtt` 

This is a super simple solution, so this is not recommand for productions

# usage and demo

    GET http://apiserver/convert-to-vtt?src=source_subtitle
    
e.g. convert ass -> vtt

    curl https://hare1039.nctu.me/convert-to-vtt?src=https://hare1039.nctu.me/sysvol/video/hare/%E8%81%96%E5%8A%8D%E9%8D%9B%E9%80%A0%E5%B8%AB/%5BYYDM-11FANS%5D%5BThe%20Sacred%20Blacksmith%5D%5B01%5D%5BBDRIP%5D%5B720P%5D%5BX264-10bit_AAC%5D%5BF63114E0%5D.ass

# dependency
- github.com/gin-gonic/gin
- `ffmpeg` installed and can be accessed by sh

