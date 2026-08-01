# THIS IS A FORK - DO NOT consider the version number as updates to other code

I'm maintaining this fork for my own use. The purpose is to merge the "BEST" parts of many of the others working on Thetis.

MW0LGE  here  https://github.com/ramdor/Thetis

Richie has "a few tweaks and changes to a ... Thetis..." I'm interested in the work he is doing on remote access.

MI0BOT  here  https://github.com/mi0bot/OpenHPSDR-Thetis

Reid has created "Radio Model" code for HERMES Lite. 

ON7OFF  here  https://github.com/ON7OFF/Thetis

Kurt has "... been working extremely hard on the Android remote software to control your Hermes Lite 2".

Each of these developers is working on different parts of the code for their own purpose. And I think them.

I have a (Hermes-Lite 2+)"http://hermeslite.com" with some personal modifications connected to a XiEGU PA125B amp. I've created this project to combine these to fit my wants and needs.  

For the full story of how this software evolved — from FlexRadio's SDR-1000 and the HPSDR/TAPR project through OpenHPSDR and Thetis to the Hermes-Lite 2 forks above — see [History](code_documentation/History.md). Developer documentation for this code base is in [code_documentation](code_documentation/README.md).

# Latest Release v2.10.3.19 15th July, 2026

# 2.10.3.19 (2026-07-15)

- Build identity: the splash screen, title bar, and About box now show **HL2 (MI0BOT/W5TSU)**. Releases from this repository are W5TSU's builds of the Hermes-Lite 2 fork, so they are not confused with MI0BOT's OpenHPSDR-Thetis releases — please report issues with these builds here, not to MI0BOT.
- Developer tooling: Doxygen configuration (`Doxyfile`) and AGENTS.md work contracts.

# 2.10.3.18 (2026-07-14)

- New SpaceBar Control "PTT/MOX Hold" option (Setup > Keyboard): hold the space bar to transmit, release to receive; works even when Focus Master hands keyboard focus to another program

# 2.10.3.17 (2026-07-13)

- N1MM spectrum: new Setup option to include/ignore the CW frequency tone shift (default off) - from official Thetis
- NuGet package updates from official Thetis
- Added code documentation under code_documentation/

# 2.10.3.15 Beta 1 (2026-06-12)

- Upgraded to release 2.10.3.15 of official Thetis
- Corrected timeout of auto tune routine
- Updated alternate RX selection form

