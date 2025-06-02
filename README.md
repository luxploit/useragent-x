# [luxploit](https://luxploit.net)/[uax](https://github.com/luxploit/useragent-x)

> [!CAUTION]
> This library is an unfinished state and is missing major bits of functionality aswell as testing as I have not yet had the time to
> do some more work on this project! **DO NOT USE FOR THE TIME BEING**.

`useragent-x` or simply `uax` is a golang 1.24+ library designed to be able to parse an extensive number of useragents, including but not limited to: Consoles, Mobile Devices, PCs and Crawlers

This library was created with inspiration from [mileusna](https://github.com/mileusna)/[useragent](https://github.com/mileusna/useragent) however as I needed an API that felt more natural to be used with my upcoming project [SpiritOnline](https://spiritonline.net) and supported a wider range of devices that I might actually see around in the wild accessing my services.

Additionally I needed to be able to distinguish between a browser that was being genuine with it's useragent and one that was impersonating another for the sake of compatbility with for example mobile devices views (the xbox comes to mind for me with its confusing UA).

All useragents tested are sourced from [danielmiessler](https://github.com/danielmiessler)'s excelent [SecLists](https://github.com/danielmiessler/SecLists/) repo under [Fuzzing/UserAgent](https://github.com/danielmiessler/SecLists/tree/master/Fuzzing/User-Agents)

If you find a device that isn't yet supported however you'd like to have added, please open a ticket so i can add it to the support list.
