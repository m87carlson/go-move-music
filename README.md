# go-move-music

I created this when Google retired Google Play Music.

This tool was ran against my Google Takeout extracted tgz archive.

Example:
```
➜  go-move-music git:(master) ✗ ls tmp
01-All_He_Has_Read.mp3                                          04-Ora_Pro_Nobis_Lucifer.mp3                                    07-Body_And_Blood.mp3                                           10-Riders_Of_Vultures.mp3
01-Blow_Your_Trumpets_Gabriel.mp3                               04-Perfect_Life.mp3                                             07-Majesty.mp3                                                  11-Ascendant_Here_On....mp3
01-First_Regret.mp3                                             04-Spoksonat.mp3                                                07-Regret_9.mp3                                                 11-La_Mantra_Mori.mp3
01-Infestissumam.mp3                                            04-Vermin.mp3                                                   08-Devil_Church.mp3                                             12-Im_a_Marionette.mp3
01-Spirit.mp3                                                   05-Amen.mp3                                                     08-Idolatrine.mp3                                               Alyans - Na Zare (At dawn).mp3
02-3_Years_Older.mp3                                            05-Ghuleh___Zombie_Queen.mp3                                    08-In_The_Absence_Ov_Light.mp3                                  Bucketheadland - Fourneau Cosmique.mp3
02-From_The_Pinnacle_To_The_Pit.mp3                             05-He_Is.mp3                                                    08-The_Sermon.mp3                                               Bucketheadland - The Five Blocks - 01 0 Block.mp3
02-Furor_Divinus.mp3                                            05-Routine.mp3                                                  08-Transience.mp3                                               Bucketheadland - The Five Blocks - 02 1 Block.mp3
02-Per_Aspera_Ad_Inferi.mp3                                     05-World_Of_Wonders.mp3                                         09-Absolution.mp3                                               Bucketheadland - The Five Blocks - 03 2 Block.mp3
02-Pillars_Of_The_South.mp3                                     06-Home_Invasion.mp3                                            09-Ancestral.mp3                                                Bucketheadland - The Five Blocks - 04 3 Block.mp3
03-Cirice.mp3                                                   06-Mummy_Dust.mp3                                               09-Depth_Of_Satans_Eyes.mp3                                     Bucketheadland - The Five Blocks - 05 4 Block.mp3
03-Hand_Cannot_Erase.mp3                                        06-The_Satanist.mp3                                             09-O_Father_O_Satan_O_Sun.mp3                                   Bucketheadland - The Five Blocks - 06 5 Block.mp3
03-Messe_Noire.mp3                                              06-The_Wind.mp3                                                 09-The_Key_And_The_Gate_(Alternative_Version_-_Bonus_Track).mp3 Буран - Воскрешение.mp3
03-Secular_Haze.mp3                                             06-Year_Zero.mp3                                                10-Deus_In_Absentia.mp3                                         
03-The_Emma.mp3                                                 07-Ben_Sahar.mp3                                                10-Happy_Returns.mp3
04-Jigolo_Har_Migiddo.mp3                                       07-Black_Sunlight.mp3                                           10-Monstrance_Clock.mp3
go-move-music git:(master) ✗ ./go-move-music -directory "tmp"
2020/11/29 18:58:27 Directory: tmp
2020/11/29 18:58:27 Gathering files
2020/11/29 18:58:27 Reading tags
 100% |████████████████████████████████████████|  [0s:0s]2020/11/29 18:58:27
Relocating songs
 100% |████████████████████████████████████████|  [0s:0s]2020/11/29 18:58:27

➜  go-move-music git:(master) ✗ tree tmp
tmp
├── Alyans\ -\ Na\ Zare\ (At\ dawn).mp3
├── Behemoth
│   └── The\ Satanist
│       ├── 01-Blow_Your_Trumpets_Gabriel.mp3
│       ├── 02-Furor_Divinus.mp3
│       ├── 03-Messe_Noire.mp3
│       ├── 04-Ora_Pro_Nobis_Lucifer.mp3
│       ├── 05-Amen.mp3
│       ├── 06-The_Satanist.mp3
│       ├── 07-Ben_Sahar.mp3
│       └── 08-In_The_Absence_Ov_Light.mp3
│       └── 09-O_Father_O_Satan_O_Sun.mp3
├── Bucketheadland\ -\ Fourneau\ Cosmique.mp3
├── Bucketheadland\ -\ The\ Five\ Blocks\ -\ 01\ 0\ Block.mp3
├── Bucketheadland\ -\ The\ Five\ Blocks\ -\ 02\ 1\ Block.mp3
├── Bucketheadland\ -\ The\ Five\ Blocks\ -\ 03\ 2\ Block.mp3
├── Bucketheadland\ -\ The\ Five\ Blocks\ -\ 04\ 3\ Block.mp3
├── Bucketheadland\ -\ The\ Five\ Blocks\ -\ 05\ 4\ Block.mp3
├── Bucketheadland\ -\ The\ Five\ Blocks\ -\ 06\ 5\ Block.mp3
├── Ghost
│   ├── Infestissumam
│   │   ├── 01-Infestissumam.mp3
│   │   ├── 02-Per_Aspera_Ad_Inferi.mp3
│   │   ├── 03-Secular_Haze.mp3
│   │   ├── 04-Jigolo_Har_Migiddo.mp3
│   │   ├── 05-Ghuleh___Zombie_Queen.mp3
│   │   ├── 06-Year_Zero.mp3
│   │   ├── 07-Body_And_Blood.mp3
│   │   ├── 08-Idolatrine.mp3
│   │   ├── 09-Depth_Of_Satans_Eyes.mp3
│   │   ├── 10-Monstrance_Clock.mp3
│   │   ├── 11-La_Mantra_Mori.mp3
│   │   └── 12-Im_a_Marionette.mp3
│   └── Meliora
│       ├── 01-Spirit.mp3
│       ├── 02-From_The_Pinnacle_To_The_Pit.mp3
│       ├── 03-Cirice.mp3
│       ├── 04-Spoksonat.mp3
│       ├── 05-He_Is.mp3
│       ├── 06-Mummy_Dust.mp3
│       ├── 07-Majesty.mp3
│       ├── 08-Devil_Church.mp3
│       ├── 09-Absolution.mp3
│       └── 10-Deus_In_Absentia.mp3
├── Steven\ Wilson
│   └── Hand.\ Cannot.\ Erase.
│       ├── 01-First_Regret.mp3
│       ├── 02-3_Years_Older.mp3
│       ├── 03-Hand_Cannot_Erase.mp3
│       ├── 04-Perfect_Life.mp3
│       ├── 05-Routine.mp3
│       ├── 06-Home_Invasion.mp3
│       ├── 07-Regret_9.mp3
│       ├── 08-Transience.mp3
│       ├── 09-Ancestral.mp3
│       ├── 10-Happy_Returns.mp3
│       └── 11-Ascendant_Here_On....mp3
├── Year\ of\ the\ Goat
│   └── The\ Unspeakable
│       ├── 01-All_He_Has_Read.mp3
│       ├── 02-Pillars_Of_The_South.mp3
│       ├── 03-The_Emma.mp3
│       ├── 04-Vermin.mp3
│       ├── 05-World_Of_Wonders.mp3
│       ├── 06-The_Wind.mp3
│       ├── 07-Black_Sunlight.mp3
│       ├── 08-The_Sermon.mp3
│       ├── 09-The_Key_And_The_Gate_(Alternative_Version_-_Bonus_Track).mp3
│       └── 10-Riders_Of_Vultures.mp3
└── \221\203\200ан\ -\ \222о\201к\200е\210ение.mp3
```
