# Byron's Ultimate AMI Deletion Package

This Golang package will help you destroy all of your AWS account's machine images. If it is an AMI, it will find it and crush it.

# Dependencies

- Go 1.5 or greater installed
- Being a dick

# Usage

Include this in your project

```
import "github.com/byronmansfield/ultimate-ami-destroyer"
```

Call the ami crush method with your amazon account id

```
ultimateAmiDestroyer.CrushAMIs(123456789)
```

Viola! No more annoying AMI's. Sit back and watch your friendly DevOps or Site Reliability Engineer get paged and loose their shit! Enjoy :)
