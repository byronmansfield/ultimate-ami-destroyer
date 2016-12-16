# Byron's Ultimate AMI Deletion Package

This Golang package will help you destroy all of your AWS account's machine images. If it is an AMI, it will find it and crush it.

# Usage

Include this in your project

```
import "ultimate-ami-destroyer"
```

Use the ami crush method with account number

```
ultimate-ami-destroyer.CrushAMIs(123456789)
```

Viola! No more annoying AMI's. Sit back and watch your friendly DevOps or Site Reliability Engineer get paged and loose their shit! Enjoy :)
